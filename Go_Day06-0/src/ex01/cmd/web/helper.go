package main

import "ex01/pkg/db/psql"

type User struct {
	Name       string
	FullAccess bool
	Thoughts   []psql.Thought
	Next       int
	Previous   int
}

const (
	host     = "localhost"
	port     = "5432"
	userdb   = "chillwavuser"
	password = "pgpwd4chillwav"
	dbname   = "chillwavdb"
)

var users = make(map[string]User)

func init() {
	users["admin"] = User{FullAccess: true, Name: "admin"}
}
