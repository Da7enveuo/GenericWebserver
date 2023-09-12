package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func (ws *WebServer) CreateDatabase() {

	// Set up a connection to the MongoDB server
	_, err := ws.Database.Exec("CREATE DATABASE IF NOT EXISTS Users")
	if err != nil {
		panic(err)
	}

	fmt.Println("Database and Table created successfully")
}

// for sign up function
func (ws *WebServer) createNewUser(user User) error {
	// Insert a document with a key of "username" and a value of the Authorization struct
	stmt, err := ws.Database.Prepare("INSERT INTO Users(id, username, createTime, active, perms, hashed_pw, salt) values (?,?,?,?,?,?)")
	if err != nil {
		fmt.Println(err)
	}
	randomNumber, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		return err
	}
	_, err = stmt.Exec(randomNumber, user.Username, time.Now().Unix(), true, user.Perms, user.HashedPassword, user.Salt)
	if err != nil {
		return err
	}

	return nil
}

func (ws *WebServer) checkUserExists(username string) (bool, error) {
	stmt, err := ws.Database.Prepare("SELECT salt, hashed_pw FROM Users WHERE username= ?")
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	d, err := stmt.Query(username)
	if err != nil {
		fmt.Println(err)
	}
	defer d.Close()
	var check bool = false
	if d.Next() {
		for d.Next() {
			err := d.Scan(&check)
			if err != nil {
				return true, err
			}
		}
	}
	return false, err
}

func (ws *WebServer) checkUserPassword(lp LoginPost) (bool, error) {
	// Find the document with the key "username" in the "profiles" collection of the "users" database
	stmt, err := ws.Database.Prepare("SELECT salt, hashed_pw FROM Users WHERE username= ?")
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	d, err := stmt.Query(lp.User)
	if err != nil {
		fmt.Println(err)
	}
	defer d.Close()
	var check bool = false

	return check, nil
}

func (ws *WebServer) checkIfDatabaseExists(databaseName string) (bool, error) {
	stmt, err := ws.Database.Prepare("SHOW DATABASES WHERE NAME = ?")
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	d, err := stmt.Query(databaseName)
	if err != nil {
		fmt.Println(err)
	}
	defer d.Close()

	if d.Next() {
		var name string
		for d.Next() {
			err = d.Scan(&name)
			if err != nil {
				return true, err
			}
			if name != databaseName {
				return true, errors.New("database names do not match")
			}
		}
	}

	return false, nil
}
