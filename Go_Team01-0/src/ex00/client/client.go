package main

import (
	"context"
	pb "ex00/gen/proto"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"strings"
	"time"
)

type Client struct {
	Conn            *grpc.ClientConn
	WareHouseClient pb.WareHouseApiClient
	CurrentInstance *Instance
	KnownInstances  []Instance
	instanceIdx     int
}

func main() {
	host := flag.String("H", "127.0.0.1", "Host")
	port := flag.String("P", "8765", "Port")
	flag.Parse()
	client := NewClient(NewInstance(*host, *port))
	for {
		err := client.Ping()
		if err != nil {
			fmt.Println(err)
			client.SwitchInstance()
			continue
		}
		fmt.Println(client.KnownInstances)
		time.Sleep(1 * time.Second)
	}

}

func (c *Client) SwitchInstance() {

	c.instanceIdx++
	fmt.Println(c.instanceIdx)
	if c.instanceIdx >= len(c.KnownInstances) {
		log.Fatalln("All instances disabled")
	}
	c.CurrentInstance = &c.KnownInstances[c.instanceIdx]
	c.Conn, _ = grpc.Dial(c.CurrentInstance.GetSocket(), grpc.WithInsecure())
	c.WareHouseClient = pb.NewWareHouseApiClient(c.Conn)
}

func (c *Client) Ping() error {
	resp, err := c.WareHouseClient.Ping(context.Background(), &pb.HeartBeatRequest{Socket: c.CurrentInstance.GetSocket()})
	if err != nil {
		return err
	}
	nodes := []Instance{}
	for i := 0; i < len(resp.KnowNodes); i++ {
		tmp := strings.Split(resp.KnowNodes[i], ":")
		nodes = append(nodes, Instance{host: tmp[0], port: tmp[1]})
	}
	c.KnownInstances = nodes
	return nil
}

func NewClient(instance *Instance) *Client {
	conn, _ := grpc.Dial(instance.GetSocket(), grpc.WithInsecure())
	return &Client{CurrentInstance: instance, Conn: conn, WareHouseClient: pb.NewWareHouseApiClient(conn)}
}

type Instance struct {
	host string
	port string
}

func (i *Instance) GetSocket() string {
	return fmt.Sprintf("%s:%s", i.host, i.port)
}

func NewInstance(host, port string) *Instance {
	return &Instance{host: host, port: port}
}
