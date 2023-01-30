package main

import (
	"context"
	pb "ex03/gen/proto"
	"ex03/psql"
	"fmt"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "chillwavuser"
	password = "pgpwd4chillwav"
	dbname   = "chillwavdb"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	db, err := psql.NewDbConnection(host, port, user, password, dbname)
	if err != nil {
		log.Fatalln(err)
	}
	err = db.CreateTable()
	if err != nil {
		log.Fatalln(err)
	}
	//data := make([]float64, 0, 100([]float64, 0, 100)
	//	var)
	var summFreq float64
	var summstd float64
	var i = 1
	for {
		client := pb.NewTransmitterApiClient(conn)
		resp, err := client.Connection(context.Background(), &pb.Request{})
		if err != nil {
			log.Fatalln(err)
		}
		summFreq += resp.Frequency
		summstd += resp.Frequency - summFreq/float64(i)
		fmt.Println(resp.Frequency, resp.Frequency-summFreq/float64(i), summFreq/float64(i), summstd/float64(i))
		i++
		if i == 50 {
			err = db.Insert(resp.SessionId, resp.Frequency, resp.UTC)
			if err != nil {
				log.Fatalln(err)
			}
			break
		}
	}
}
