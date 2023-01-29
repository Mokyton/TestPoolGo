package main

import (
	"context"
	pb "ex00/gen/proto"
	"fmt"
	uuid2 "github.com/satori/go.uuid"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"math/rand"
	"net"
	"os"
)

type TransmitterApiServer struct {
	pb.UnimplementedTransmitterApiServer
}

func (s *TransmitterApiServer) Connection(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	res := &pb.Response{}
	res.SessionId = uuid2.NewV4().String()
	res.Frequency = rand.Float64()*-10 + 10
	std := rand.Float64()*0.3 + 1.5
	res.UTC = timestamppb.Now()
	_, err := fmt.Fprintln(os.Stdout, res.UTC, res.SessionId, res.Frequency, std, "\n")
	if err != nil {
		return nil, err
	}
	return res, nil
}

func main() {
	listner, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterTransmitterApiServer(grpcServer, &TransmitterApiServer{})

	err = grpcServer.Serve(listner)
	if err != nil {
		log.Println(err)
	}
}
