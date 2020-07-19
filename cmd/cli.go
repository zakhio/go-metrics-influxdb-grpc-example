package main

import (
	"context"
	"fmt"
	"github.com/zakhio/go-metrics-influxdb-grpc-example/proto"
	"google.golang.org/grpc"
)

const (
	port = ":50053"
)

func main() {
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := proto.NewEchoServiceClient(conn)

	for {
		var msg string
		fmt.Print("> ")
		_, err := fmt.Scanf("%s\n", &msg)
		if err != nil {
			panic(err)
		}

		res, err := client.Echo(context.Background(), &proto.EchoRequest{Message: msg})
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v\n", res.Message)
	}
}
