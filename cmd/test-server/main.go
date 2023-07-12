package main

import (
	"context"
	"fmt"
	"ghz-test-sfixed64/protos"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type TestService struct{}

func (s TestService) Echo(ctx context.Context, request *protos.TestRequest) (*protos.TestResponse, error) {
	fmt.Println("Id: ", request.Id)
	return &protos.TestResponse{Id: request.Id}, nil
}

func main() {
	srv := grpc.NewServer()
	protos.RegisterTestServiceServer(srv, &TestService{})
	lis, err := net.Listen("tcp", ":9009")
	if err != nil {
		panic(err)
	}
	errs := make(chan error, 1)
	go func() { errs <- srv.Serve(lis) }()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	for {
		select {
		case e := <-errs:
			if err != nil {
				fmt.Println("Fatal error...", e)
				return
			}
		case <-sig:
			fmt.Println("Graceful shutdown...")
			srv.GracefulStop()
			return
		}
	}
}
