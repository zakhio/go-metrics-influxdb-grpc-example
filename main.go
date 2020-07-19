package main

import (
	"github.com/rcrowley/go-metrics"
	"github.com/zakhio/go-metrics-influxdb"
	"github.com/zakhio/go-metrics-influxdb-grpc-example/config/influxdb"
	"github.com/zakhio/go-metrics-influxdb-grpc-example/handler"
	"github.com/zakhio/go-metrics-influxdb-grpc-example/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50053"
)

func main() {
	influxDBConfig := influxdb.ParseConfig("influxdb-config.yaml")
	go reporter.InfluxDB(metrics.DefaultRegistry,
		influxDBConfig.Interval,
		influxDBConfig.URL,
		influxDBConfig.OrganizationId,
		influxDBConfig.BucketId,
		influxDBConfig.Measurement,
		influxDBConfig.Token,
		influxDBConfig.AlignTimestamps,
	)

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterEchoServiceServer(grpcServer, handler.NewEchoService())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
