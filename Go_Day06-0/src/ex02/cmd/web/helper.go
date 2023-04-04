package main

import (
	"ex01/pkg/db/psql"
	"fmt"
	"net"
	"net/http"
	"sync/atomic"
)

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

var (
	users       = make(map[string]User)
	connections = Connection{Counte: atomic.Int32{}}
)

type Connection struct {
	Counte atomic.Int32
}

func (c *Connection) OnstateChange(conn net.Conn, state http.ConnState) {
	switch state {
	case http.StateNew:
		c.Add(1)
	case http.StateHijacked, http.StateClosed:
		c.Add(-1)
	}
	fmt.Println(c.Counte.Load())
}

func (c *Connection) Add(num int32) {
	c.Counte.Add(num)
}

func init() {
	users["admin"] = User{FullAccess: true, Name: "admin"}
}
