package main

import (
	"net/http"

	"database/sql"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
)

type Context struct {
	DomainName    string
	HTTPSPEMFile  string
	HTTPSKeyFile  string
	RedisIP       string
	RedisCert     string
	MySQLIP       string
	MySQLUser     string // and this one
	MySQLPassword string // please secure this
}

type WebServer struct {
	Router   *http.ServeMux
	Cache    *Cache
	Context  *Context
	Database *sql.DB
}

// session manager structs
type SessionMgr struct {
	Sessions map[string]Sesh
}

type Sesh struct {
	Flagged      bool
	RequestCount int
}

// cache structs
type Cache struct {
	RedisInstance *redis.Client
}

// user auth struct
type LoginPost struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

// the Database entry will have a key of random generated uuid and value of this struct
type User struct {
	Username       string
	HashedPassword string
	Salt           string
	Perms          []bool
}
