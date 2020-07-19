package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/zakhio/go-metrics-influxdb-grpc-example/proto"
	"google.golang.org/grpc"
)

const (
	address = ":50053"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := proto.NewEchoServiceClient(conn)

	for {
		var msg string
		fmt.Print("> ")
		_, err := fmt.Scanf("%s\n", &msg)
		if err != nil {
			log.Fatalf("could not scanf: %v", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		res, err := client.Echo(ctx, &proto.EchoRequest{Message: msg})
		if err != nil {
			log.Fatalf("could not echo: %v", err)
		}
		log.Printf("%v\n", res.Message)
	}
}
