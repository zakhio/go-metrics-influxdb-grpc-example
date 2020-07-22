package handler

import (
	"context"
	"fmt"
	"sync/atomic"

	"github.com/rcrowley/go-metrics"
	"github.com/zakhio/go-metrics-influxdb-grpc-example/proto"
)

var (
	// A uniform sample using Vitter's Algorithm R.
	// <http://www.cs.umd.edu/~samir/498/vitter.pdf>
	messageLengthHistogram = metrics.GetOrRegisterHistogram("message_length", nil, metrics.NewUniformSample(50))
	messageCounter         = metrics.GetOrRegisterCounter("message", nil)
)

type echoServer struct {
	proto.UnimplementedEchoServiceServer
	counter uint64
}

//Echo sends back message from the request and its number
func (s *echoServer) Echo(_ context.Context, req *proto.EchoRequest) (*proto.EchoResponse, error) {
	messageLengthHistogram.Update(int64(len(req.Message)))
	messageCounter.Inc(1)

	count := atomic.AddUint64(&s.counter, 1)

	return &proto.EchoResponse{Message: fmt.Sprintf("Message #%d: %v", count, req.Message)}, nil
}

//NewEchoService instantiate new echo server object
func NewEchoService() proto.EchoServiceServer {
	server := &echoServer{}
	return server
}
