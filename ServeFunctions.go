package main

import (
	"io"
	"net/http"
)

// serving the login page
func serveLogin(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "login.jsx")

}

// serving the dashboard page
func serveDashboard(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "dashboard.jsx")
}

// serving settings page
func serveSettings(w http.ResponseWriter, r *http.Request) {
	// need to check if there is a user name invovled
	http.ServeFile(w, r, "settings.jsx")
}

func servePing(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Pong")
}

func serveSignup(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "signup.jsx")
}

func serveErrorPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "error.html")
}

/*
func buildHandler() {

}
*/
