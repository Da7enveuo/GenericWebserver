package main

func (ws *WebServer) setupDB() {
	// change accordingly

	check, err := ws.checkIfDatabaseExists("Authentication")
	if err != nil {
		panic(err)
	}
	if !check {
		// check this shit
		ws.CreateDatabase()
	}

}
