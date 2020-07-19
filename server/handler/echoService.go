package handler

import (
	"context"
	"fmt"
	"sync/atomic"

	"github.com/rcrowley/go-metrics"
	"github.com/zakhio/go-metrics-influxdb-grpc-example/proto"
)

var (
	echoCounter = metrics.GetOrRegisterCounter("echo", nil)
)

type echoServer struct {
	proto.UnimplementedEchoServiceServer
	counter uint64
}

func (s *echoServer) Echo(ctx context.Context, req *proto.EchoRequest) (*proto.EchoResponse, error) {
	echoCounter.Inc(1)
	c := atomic.AddUint64(&s.counter, 1)

	return &proto.EchoResponse{Message: fmt.Sprintf("Message #%d: %v", c, req.Message)}, nil
}

func NewEchoService() proto.EchoServiceServer {
	server := &echoServer{}
	return server
}
