package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

func isAuthenticated(nextEndpoint http.HandlerFunc, ws *WebServer) http.HandlerFunc {
	parentContext, cancel := context.WithTimeout(context.Background(), 500*time.Second)
	defer cancel()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionId, err := r.Cookie("SessionID")
		if err != nil {
			// if session cookie is not there, redirect them to login
			serveLogin(w, r.WithContext(parentContext))
		}
		// need some input sanitization here before we get the session info
		sess := sanitizeInput(string(sessionId.Value))
		sesh, err := ws.getSessionInfo(sess)
		if err != nil {
			panic(err)
		}
		// what the fuck we pushing here? should push new updates
		ws.pushValue(string(sessionId.Value), *sesh)
		authOrNo := ws.checkCache(sessionId.Value)
		if !authOrNo {
			serveLogin(w, r.WithContext(parentContext))
		} else {
			nextEndpoint(w, r.WithContext(parentContext))
		}
	})
}

func ParameterCheck(nextEndpoint http.HandlerFunc, ws *WebServer) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			// assumign that the params are formatted like <url stuff>&key=val&key=val
			// so this breaks it to look like: key=val, key=val, key=val.
			array_param := strings.SplitN(r.URL.Path, "&", -1)
			var params map[string]interface{}
			for _, kv := range array_param {
				param, err := url.QueryUnescape(kv)
				if err != nil {
					log.Fatal(err)
					serveErrorPage(w, r)

				}
				// have to do checking on these user controlled params here, just regex check her

				kv_pair := strings.SplitN(param, "=", -1)
				// saying its a empty map, but it should not be.

				params[kv_pair[0]] = kv_pair[1]
			}

		default:
			serveErrorPage(w, r)
		}
	})
}
func LoginPageCheck(nextEndpoint http.HandlerFunc, ws *WebServer) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			http.ServeFile(w, r, "login.jsx")
		} else if r.Method == "POST" {
			var body LoginPost
			bd, err := ioutil.ReadAll(r.Body)
			if err != nil {
				panic(err)
			}
			err = json.Unmarshal(bd, &body)
			if err != nil {
				panic(err)
			}
			check, err := ws.checkUserPassword(body)
			if err != nil {
				panic(err)
			}
			if check {
				sessioninfo := Sesh{Flagged: false, RequestCount: 1}
				err = ws.createSession(sessioninfo)
				if err != nil {
					panic(err)
				}
				serveDashboard(w, r)
			}
		}
	})
}

func isAuthorized(nextEndpoint http.HandlerFunc, ws *WebServer) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if the user is authenticated then the session information will be in the request context, per the function above and the function chain described in routes.
		parentctx := r.Context()
		sessionInfo := parentctx.Value("SessionInfo")
		// unsure what type session info will be
		fmt.Println(sessionInfo)
		// create function that will take in the endpoint requested and gather lowest permission level, then compare with the session info provided in context from isauthenticated.
	})
}

func signupWrapper(nextEndpoint http.HandlerFunc, ws *WebServer) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST" {
			var body LoginPost
			bd, err := ioutil.ReadAll(r.Body)
			if err != nil {
				panic(err)
			}
			err = json.Unmarshal(bd, &body)
			if err != nil {
				panic(err)
			}
			uname := sanitizeInput(body.User)
			pw, salt := createHashedPassword(body.Password)
			dbE := User{
				Username:       uname,
				HashedPassword: pw,
				Salt:           salt,
				Perms:          []bool{},
			}
			err = ws.createNewUser(dbE)
			if err != nil {
				serveErrorPage(w, r)
			}
			serveDashboard(w, r)

		} else if r.Method == "GET" {
			serveSignup(w, r)
		} else {
			serveErrorPage(w, r)
		}
		// if the user is authenticated then the session information will be in the request context, per the function above and the function chain described in routes.
		parentctx := r.Context()
		sessionInfo := parentctx.Value("SessionInfo")
		// unsure what type session info will be
		fmt.Println(sessionInfo)
		// create function that will take in the endpoint requested and gather lowest permission level, then compare with the session info provided in context from isauthenticated.
	})
}

func sanitizeInput(input string) string {
	// Remove leading and trailing whitespace
	input = strings.TrimSpace(input)

	// Replace any special characters with underscores
	re := regexp.MustCompile("[^a-zA-Z0-9_]")
	input = re.ReplaceAllString(input, "_")

	return input
}
func signoutWrapper(nextEndpoint http.HandlerFunc, ws *WebServer) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			sessionId, err := r.Cookie("SessionID")
			if err != nil {
				panic(err)
			}
			err = ws.deprecateSession(sessionId.Value)
			if err != nil {
				serveSignup(w, r)
			}
			serveLogin(w, r)

		} else {
			serveErrorPage(w, r)
		}
	})
}
