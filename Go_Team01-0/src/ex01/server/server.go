package main

import (
	"context"
	"errors"
	pb "ex01/gen/proto"
	"fmt"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/anypb"
	"net"
	"sync"
)

var Nodes []Instance
var storage map[string]map[any]any

type WareHouseApi struct {
	pb.UnimplementedWareHouseApiServer
}

func main() {
	storage = make(map[string]map[any]any)
	var wg sync.WaitGroup
	Nodes = []Instance{{host: "127.0.0.1", port: "8765"}, {host: "127.0.0.1", port: "9876"}}
	listeners := make([]net.Listener, len(Nodes), len(Nodes))
	serves := make([]*grpc.Server, len(Nodes), len(Nodes))

	for i := 0; i < len(listeners); i++ {
		listeners[i], _ = net.Listen("tcp", Nodes[i].GetSocket())
	}

	srv := &WareHouseApi{}

	for i := 0; i < len(serves); i++ {
		wg.Add(1)
		serves[i] = grpc.NewServer()
		pb.RegisterWareHouseApiServer(serves[i], srv)
		go func(i int) {
			defer wg.Done()
			serves[i].Serve(listeners[i])
		}(i)
	}
	for i := 0; i < len(Nodes); i++ {
		fmt.Println("WareHouse starts at", Nodes[i].GetSocket())
	}

	wg.Wait()
}

func (w *WareHouseApi) Ping(ctx context.Context, req *pb.HeartBeatRequest) (*pb.HeartBeatResponse, error) {
	nodes := make([]string, len(Nodes))
	for i := 0; i < len(nodes); i++ {
		nodes[i] = Nodes[i].GetSocket()
	}
	return &pb.HeartBeatResponse{CurrentInstance: req.Socket, KnowNodes: nodes}, nil
}

func (w *WareHouseApi) Set(ctx context.Context, req *pb.SetRequest) (*pb.SetResponse, error) {
	cluster, ok := storage[req.Cluster]
	if !ok {
		init := make(map[any]any)
		v, _ := proto.Marshal(req.Value)
		init[req.Key.String()] = string(v)
		storage[req.Cluster] = init
	} else {
		v, _ := proto.Marshal(req.Value)
		cluster[req.Key.String()] = string(v)
	}
	return &pb.SetResponse{Success: true}, nil
}

func (w *WareHouseApi) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	cluster, ok := storage[req.Cluster]
	if !ok {
		return &pb.DeleteResponse{Success: false}, errors.New("Cluster doesn't exist ")
	}
	delete(cluster, req.Key.String())
	return &pb.DeleteResponse{Success: true}, nil
}

func (w *WareHouseApi) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	cluster, ok := storage[req.Cluster]
	if !ok {
		return nil, errors.New("Cluster doesn't exist ")
	}
	v, ok := cluster[req.Key.String()]
	if !ok {
		return nil, errors.New("Key doesn't exist ")
	}
	value := []byte(v.(string))
	return &pb.GetResponse{Value: &anypb.Any{Value: value}}, nil
}
