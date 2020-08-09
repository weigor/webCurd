package main

import (
	"awesomeProject3/db"
	order_proto "awesomeProject3/rpc/server/proto"
	"awesomeProject3/rpc/server/service"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func main() {
	fmt.Println("grpc 准备开启--")
	err := db.InitMySQL()
	if err != nil {
		panic(err)
	}
	//关闭数据库连接
	defer db.Close()
	rpcServer := grpc.NewServer()
	order_proto.RegisterOrderServiceServer(rpcServer, new(service.Orderservice))

	listener, _ := net.Listen("tcp", ":8022")

	if err := rpcServer.Serve(listener); err != nil {
		fmt.Println(err)
	}
	fmt.Println(listener)
	fmt.Println("grpc 开启完成!")
}
