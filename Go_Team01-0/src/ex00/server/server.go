package main

import (
	"context"
	pb "ex00/gen/proto"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"time"
)

var Nodes []Instance

type WareHouseApi struct {
	pb.UnimplementedWareHouseApiServer
}

func main() {

	Nodes = []Instance{{host: "127.01.0.1", port: "8765"}, {host: "127.01.0.1", port: "9876"}}
	listeners := make([]net.Listener, len(Nodes), len(Nodes))

	serves := make([]*grpc.Server, len(Nodes), len(Nodes))

	for i := 0; i < len(listeners); i++ {
		listeners[i], _ = net.Listen("tcp", Nodes[i].GetSocket())
	}

	t := time.After(10 * time.Second)
	srv := &WareHouseApi{}
	for i := 0; i < len(serves); i++ {
		serves[i] = grpc.NewServer()
		pb.RegisterWareHouseApiServer(serves[i], srv)
		go serves[i].Serve(listeners[i])
	}

	for i := 0; i < len(Nodes); i++ {
		fmt.Println("WareHouse starts at", Nodes[i].GetSocket())
	}

	select {
	case <-t:
		serves[0].Stop()
	}
}

func (w *WareHouseApi) Ping(ctx context.Context, req *pb.HeartBeatRequest) (*pb.HeartBeatResponse, error) {
	nodes := make([]string, len(Nodes))
	for i := 0; i < len(nodes); i++ {
		nodes[i] = Nodes[i].GetSocket()
	}
	return &pb.HeartBeatResponse{CurrentInstance: req.Socket, KnowNodes: nodes}, nil
}
