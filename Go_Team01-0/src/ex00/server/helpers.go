package main

import "fmt"

type Instance struct {
	host string
	port string
}

func (i *Instance) GetSocket() string {
	return fmt.Sprintf("%s:%s", i.host, i.port)
}
