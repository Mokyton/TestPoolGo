package main

import (
	"ex02/psql"
	_ "github.com/lib/pq"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "habrpguser"
	password = "pgpwd4habr"
	dbname   = "habrdb"
)

func main() {
	conn, err := psql.NewDbConnection(host, port, user, password, dbname)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.DB.Close()
	err = conn.CreateTable()
	err = conn.Insert("helpmeeeeeeee", 0.2168465, timestamppb.Timestamp{Nanos: 865})
	if err != nil {
		log.Fatal(err)
	}
}
