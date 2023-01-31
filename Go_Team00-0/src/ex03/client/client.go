package main

import (
	"context"
	pb "ex03/gen/proto"
	"ex03/psql"
	"flag"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
	"math"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "chillwavuser"
	password = "pgpwd4chillwav"
	dbname   = "chillwavdb"
)

func main() {
	k := flag.Float64("k", 10, "anomaly coefficient")
	flag.Parse()
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
	data := make([]float64, 0, 100)
	var i = 0
	for {
		client := pb.NewTransmitterApiClient(conn)
		resp, err := client.Connection(context.Background(), &pb.Request{})
		if err != nil {
			log.Fatalln(err)
		}
		data = append(data, resp.Frequency)
		sigma := sigma(data)
		tNegative := resp.Frequency - 2*sigma
		tPozitive := resp.Frequency + 2*sigma
		if sigma**k > tPozitive {
			err = db.Insert(resp.SessionId, tPozitive, resp.UTC.AsTime().String())
			if err != nil {
				log.Fatalln(err)
			}
		}
		if sigma*-*k < tNegative {
			err = db.Insert(resp.SessionId, tNegative, resp.UTC.AsTime().String())
			if err != nil {
				log.Fatalln(err)
			}
		}
		i++
	}
}

func sigma(data []float64) float64 {
	m := mean(data)
	var sum float64
	for _, v := range data {
		sum += math.Pow(v-m, 2)

	}
	return math.Sqrt(sum / float64(len(data)))
}

func mean(nums []float64) float64 {
	var sum float64
	for _, v := range nums {
		sum += v
	}
	return sum / float64(len(nums))
}
