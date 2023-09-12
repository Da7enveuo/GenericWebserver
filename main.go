package main

func main() {
	// will likely need to use some type of dependency injection tool
	// we will need the domain that we want to serve from, the topic address within kafak, the ip address of the cache, the certificate of the cache, and the id of the pod/cluster created
	// this will return a pointer to our context structure, which will be used throughout the program
	//               domain                https pem                         https key                 redis ip           redis cert                      mongo ip		mongo user	  mongo password
	con := &Context{"localserver.localhost", "/etc/ssl/certs/webserver.pem", "/etc/ssl/certs/webserver.key", "127.0.0.1:9000", "/etc/ssl/certs/certs/redis.crt", "127.0.0.1:27017", "mongoUser", "mongoPassword"}
	// need to grab ip address of kafka cluster to send producer data
	// need to get the database address

	webserver, err := constructor(con)

	if err != nil {
		panic(err)
	}
	// this configures the webservers web routes
	webserver.routes()
	webserver.setupDB()
	// start web server
	mErr := webserver.start()
	// test webserver
	err = webserver.pingWebServer()
	if err != nil {
		panic(err)
	}
	// test cache
	err = webserver.pingCache()
	if err != nil {
		panic(err)
	}
	// will need to determine if we want the analysis on a separate server

	/*
		Server 1: Redis Cache (Internal Only)
					Holds session cached information
		Server 2: Mongo DB (Internal Only)
					holds users and passwords
					holds email analysis data
		Server 3: Email Analysis Web Server (Internal/External)
					displays data from Mongo Cache, uses redis and mongo for user authentication/authorization/sessionmanagement
		Server 4: Email Connector (Internal Only or Only allowed to talk to microsoft)
					pulls emails that come in to analyze them and correlate any further potential matches
		Server 5: Analysis Server (External/Internal [OSINT Analysis])
					Does osint analysis on IoCs found in the email body/header/attachements
						reconstruct spf path to determine if the email potentially traveled to a malicious node
						spf and dkim and dmarc checking

		Server 6: NLP Sentiment Classification (GPT3) (Internal Only)
					Classifies whether the email text body appears to be malicious/suspicious

		* server 1-5 may be able to be hosted on the same server w/docker, but NLP sentiment classification should be on a separate one
	*/

	if mErr != nil {
		panic(mErr)
	}

}
