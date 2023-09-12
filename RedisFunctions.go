package main

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
)

func (ws *WebServer) createSession(Ses Sesh) error {
	new_session := uuid.New().String() + uuid.New().String()
	// create random uuid's, will check if in cache then cement the creation within redis
	_, err := ws.Cache.RedisInstance.Get(new_session).Result()
	if err != nil {
		return err
	} else {
		// sessioninformation is going to be context.Context and hold useful session information within. Will be adding onto to file constantly
		// default cache time of 5 minutes

		err = ws.Cache.RedisInstance.Set(new_session, Ses, 600).Err()
		if err != nil {
			return err
		}
		return nil
	}
}

// would be for updating users in the cache
func (ws *WebServer) pushValue(sesid string, ses Sesh) error {
	// Convert the Authorization struct to JSON
	sesJson, err := json.Marshal(ses)
	if err != nil {
		return err
	}

	// Push the value to the Redis cache
	err = ws.Cache.RedisInstance.LPush(sesid, sesJson).Err()
	if err != nil {
		return err
	}

	return nil
}

func (ws *WebServer) checkCache(sessionid string) bool {
	// Check if the key exists in the Redis cache
	exists, err := ws.Cache.RedisInstance.Exists(sessionid).Result()
	if err != nil {
		log.Fatal(err)
	}

	return exists == 1
}

// this would be when a user signs out or
func (ws *WebServer) deprecateSession(sessionid string) error {
	// first we check if the session provided exists
	err := ws.Cache.RedisInstance.Del(sessionid).Err()
	if err != nil {
		return err
	}
	return nil

}

func (ws *WebServer) getSessionInfo(sessionid string) (*Sesh, error) {
	// Execute a Redis command to retrieve the JSON structure from the cache
	jsonStr, err := ws.Cache.RedisInstance.Get(sessionid).Result()
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON structure into a Sesh struct
	var sesh Sesh
	if err := json.Unmarshal([]byte(jsonStr), &sesh); err != nil {
		return nil, err
	}
	return &sesh, nil
}
