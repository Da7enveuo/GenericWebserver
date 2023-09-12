package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-redis/redis"
)

func init() {
	log.Printf("Starting Server on Port 443")
}

// we will probably pull them as environment variables, want some easy way to store keys created, may just end up pushing to kafka queue

func constructor(Context *Context) (*WebServer, error) {
	//RouteMap := map[string] http.HandlerFunc {
	//	"/login":
	//}
	db, err := sql.Open("mysql", "<user>:<password>@/<databasename>")
	if err != nil {
		panic(err.Error())
	}

	// close database after all work is done
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	ws := &WebServer{
		// routing requests between different endpoints
		Router: http.NewServeMux(),
		// for caching credentials, session information, etc
		Cache: &Cache{
			RedisInstance: redis.NewClient(&redis.Options{
				// The addr below will be dynmaic based on whatever the internal cluster address is for redis cache
				Addr:     Context.RedisIP,
				Password: Context.RedisCert,
				DB:       10,
			})},
		// This will be things to assist with connecting to the different services in the network
		Context: Context,
		// gotta figure this out for redis cache and database, do I want to instantiate every request a new client, do I want to instantiate once and it live forever? Do I want to implement more complex solution that will utilize a certain number of connections to db efficently? YES
		Database: db,
		// Will declare the address and the handler later for kafka
		// Will declare handler for mongo later, likely will be an address in aws
		/*
			Database:
		*/
	}

	return ws, nil
}
