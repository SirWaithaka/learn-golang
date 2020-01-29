package main

import (
	"context"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/sirwaithaka/gorpc/proto"
)

type server struct{}

func main() {
	// create tcp listener on port 4040
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	// create a grpc server
	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	// serve grpc server on the tcp port
	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}

// define Add function that implements the protobuf interface
func (s *server) Add(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a + b
	return &proto.Response{Result: result}, nil

}

// define a
func (s *server) Multiply(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a * b
	return &proto.Response{Result: result}, nil

}
