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
	client := pb.NewTransmitterApiClient(conn)

	resp, err := client.Connection(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp)
}
