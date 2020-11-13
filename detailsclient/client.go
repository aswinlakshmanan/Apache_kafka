package main

import (
	"context"
	"fmt"
	"log"

	"./proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connectiton failed: %v, err")
	}
	defer conn.Close()

	c := proto.NewDetailsserviceClient(conn)

	req := proto.detailsrequest{
		details: &proto.details{
			username: "aswin1998",
			name:     "aswin lakshmanan",
		},
	}
	res, err := c.info(context.Background(), &req)
	if err != nil {
		log.Fatalf("request failed: %v", err)
	}
	fmt.Println(res.Result)

}
