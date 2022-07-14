package client

import (
	"context"
	"log"

	"github.com/wdfky/grpcdemo/pb"
	"google.golang.org/grpc"
)

func Init() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewGreeterClient(conn)

	ret, err := client.SayHello(context.TODO(), &pb.HelloRequest{
		Name: "6666",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(ret.GetMessage())
}
