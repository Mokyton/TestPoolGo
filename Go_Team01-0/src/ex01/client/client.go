package main

import (
	"bufio"
	"context"
	"encoding/json"
	pb "ex01/gen/proto"
	"flag"
	"fmt"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/anypb"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
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
	err := client.Ping()
	if err != nil {
		client.SwitchInstance()
	}

	fmt.Println("Connected to a database of Warehouse 13 at", client.CurrentInstance.GetSocket())
	fmt.Println("Known nodes:")
	for i := 0; i < len(client.KnownInstances); i++ {
		fmt.Println(client.KnownInstances[i].GetSocket())
	}
	kill := make(chan os.Signal, 1)
	signal.Notify(kill, syscall.SIGKILL, os.Interrupt, syscall.SIGTERM)

	scann := bufio.NewScanner(os.Stdin)
	for scann.Scan() {
		select {
		case <-kill:
			os.Exit(0)
		default:
			statement := strings.Fields(scann.Text())
			switch statement[0] {
			case "SET":
				if len(statement) != 4 {
					break
				}
				key, _ := json.Marshal(statement[2])
				value, _ := json.Marshal(statement[3])
				resp, err := client.WareHouseClient.Set(context.Background(),
					&pb.SetRequest{
						Cluster: statement[1],
						Key:     &anypb.Any{Value: key},
						Value:   &anypb.Any{Value: value},
					})
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(resp.String())
				}
			case "GET":
				if len(statement) != 3 {
					break
				}
				key, _ := json.Marshal(statement[2])
				resp, err := client.WareHouseClient.Get(context.Background(),
					&pb.GetRequest{
						Cluster: statement[1],
						Key:     &anypb.Any{Value: key},
					})
				if err != nil {
					fmt.Println(err)
				} else {
					v, _ := proto.Marshal(resp.Value)
					fmt.Println(string(v))
				}
			case "DELETE":
				if len(statement) != 3 {
					break
				}
				key, _ := json.Marshal(statement[2])
				resp, err := client.WareHouseClient.Delete(context.Background(),
					&pb.DeleteRequest{
						Cluster: statement[1],
						Key:     &anypb.Any{Value: key},
					})
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(resp.String())
				}
			case "/q":
				os.Exit(0)
			default:
				fmt.Println("Wrong request")
			}
		}
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
