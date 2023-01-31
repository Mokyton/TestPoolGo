package main

import (
	"context"
	pb "ex00/gen/proto"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"math"
	"os"
)

func main() {
	k := flag.Float64("k", 10, "anomaly coefficient")
	flag.Parse()
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
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
			fmt.Fprintln(os.Stdout, tPozitive)
		}
		if sigma*-*k < tNegative {
			fmt.Fprintln(os.Stdout, tNegative)
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
