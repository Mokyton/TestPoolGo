package main

import (
	"context"
	pb "ex00/gen/proto"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	//data := make([]float64, 0, 100)
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
			break
		}
	}
}
