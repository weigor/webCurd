package main

import (
	s "awesomeProject3/rpc/protobuf"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8022", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	client := s.NewOrderServiceClient(conn)
	req := &s.GetOrderRequest{Id: 33}
	order, err := client.GetOrderList(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(order.Id)
	fmt.Println(order.FileUrl)
	fmt.Println(order.Status)
	fmt.Println(order.Amount)
	fmt.Println(order.UserName)
	fmt.Println(order.OrderNo)

}
