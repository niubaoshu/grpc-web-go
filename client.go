package main

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	pb "test/testgrpc/helloworld"
)

const (
	address     = "localhost:9091"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)

	r3, err := c.SayRepeatHello(context.Background(), &pb.RepeatHelloRequest{Name: name, Count: 5})
	if err != nil {
		log.Printf("could not greet: %v", err)
	} else {
		for {
			r, err := r3.Recv()
			if err != nil {
				log.Printf("could not greet: %v", err)
				break
			} else {
				log.Printf(r.Message)
			}
		}
	}

	rd, err := c.SayHelloAfterDelay(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Printf("could not greet: %v", err)
	} else {
		log.Printf(rd.Message)
	}

}
