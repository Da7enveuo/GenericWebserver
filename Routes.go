package main

import "net/http"

// used to iniatlize the routes for the web server
func (ws *WebServer) routes() {
	// we pass in the web server struct to the isAuthenticated and isAuthorized middleware methods because we cannot inherit via the typical (ws *WebServer) way, we probably could using just the handle function but will adjust later if necessary
	// obviously the login function, will need to be a post with a specified
	ws.Router.HandleFunc("/login", authAbs(LoginPageCheck(serveLogin, ws), ws))
	// specialized dashboard based on the user
	ws.Router.HandleFunc("/dashboard", authAbs(serveDashboard, ws))

	// this is for the users settings
	ws.Router.HandleFunc("/settings", authAbs(serveSettings, ws))
	// this is to test that the server is active externally
	ws.Router.HandleFunc("/ping", servePing)

	ws.Router.HandleFunc("/signup", signupWrapper(serveSignup, ws))
	ws.Router.HandleFunc("/signout", signoutWrapper(serveLogin, ws))
}

func authAbs(endpoint http.HandlerFunc, ws *WebServer) http.HandlerFunc {
	return isAuthenticated(isAuthorized(endpoint, ws), ws)
}

func (ws *WebServer) addCustomRoutes() {
	// here I want to read a custom file that points to html/js/woff/css files to serve and add them into the web routes
}

/* example configuration file to create endpoints
{
	web_endpoint: "/api/users",
	file_location: "either ip for nodejs shit or file for regular html",
	authenticated: True,
	// should we also have a data location for the params, to define behavior behind them and such? like authenticate or some shit?
	params: [
		{
			username: {DBParam bool, CacheParam bool, datatype: bool|string|int}

		},
	]
}


*/
