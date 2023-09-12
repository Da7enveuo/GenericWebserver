package main

import (
	"errors"
	"fmt"
	"net/http"
)

// The constructor initalizes several variables for the code to function (redis cluster IP, domain name, local ip, etc..)
type ContextCreation interface {
	constructor() (*WebServer, error)
}

type WebServerComponents interface {
	// routes create the endpoints for the web server.
	routes()
	// start the web server
	start() error
	// This pings to ensure that a webserver is working
	pingWebServer() error
	// This pings the redis cache to ensure a good connection
	pingCache() error
	// this creates a session
	createSession() error
	// this checks if a user is in the data base
	checkIfUserExists() error
	// this deprecates a session
	deprecateSession() error
	// this gracefully shuts the server
	shutdown() error

	pushValue() error
	createRedisEntry() error
	checkCache() error

	CreateDatabase()
	createNewUser() error
	checkUserPassword() error
	QueryMongoDB(interface{}) (interface{}, error)
	setupDB() error
}

func (ws *WebServer) pingWebServer() error {
	// must get the domain name from the docker image/kubernetes image/aws keyvault
	resp, err := http.Get(fmt.Sprintf("https://%v/ping", ws.Context.DomainName))
	if err != nil {
		return err
	} else if resp.StatusCode != 200 {
		return errors.New("malformed status code")
	}
	return nil
}

func (ws *WebServer) start() error {
	err := http.ListenAndServeTLS(ws.Context.DomainName, ws.Context.HTTPSPEMFile, ws.Context.HTTPSKeyFile, ws.Router)
	return err
}

func (ws *WebServer) pingCache() error {
	_, err := ws.Cache.RedisInstance.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
