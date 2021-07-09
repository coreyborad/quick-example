package main

import (
	"context"
	"fmt"
	"go-grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	conn, _ := grpc.Dial("localhost:9090", grpc.WithInsecure())
	client := proto.NewBasicServiceClient(conn)
	for {
		result, err := client.GetTimeAndName(context.Background(), &proto.ServerInfo{
			Name: "CLIENT",
			Time: timestamppb.Now(),
		})
		if err != nil {
			fmt.Println(err, "ER")
		}
		fmt.Println(result)
	}

}
