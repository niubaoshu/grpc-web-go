package main

import (
	"context"
	"log"
	"net"
	"strconv"
	"time"

	pb "./helloworld"
	"google.golang.org/grpc"
)

const (
	port = ":9091"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) SayRepeatHello(in *pb.RepeatHelloRequest, sender pb.Greeter_SayRepeatHelloServer) error {
	for i := int32(0); i < in.Count; i++ {
		err := sender.Send(&pb.HelloReply{Message: in.Name + strconv.FormatInt(int64(i), 10)})
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *server) SayHelloAfterDelay(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	time.Sleep(time.Second * 5)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	//reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
