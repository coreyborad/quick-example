package main

import (
	"fmt"
	"net"

	"go-grpc/handler"
	"go-grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	fmt.Println("START")
	GRPCServe(":9090")
}

// GRPC GRPC
type GRPC struct {
	grpc *grpc.Server

	Addr string
}

// ListenAndServe ListenAndServe
func (srv *GRPC) ListenAndServe() error {
	addr := srv.Addr
	if addr == "" {
		addr = ":9090"
	}
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	proto.RegisterBasicServiceServer(srv.grpc, handler.CreateExampleHandler())

	reflection.Register(srv.grpc)

	return srv.grpc.Serve(lis)
}

// GRPCServe GRPCServe
func GRPCServe(addr string) error {
	server := &GRPC{
		grpc: grpc.NewServer(),
		Addr: addr,
	}

	return server.ListenAndServe()
}
