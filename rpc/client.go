package main


import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

//请求参数
type ComputRequset struct {
	A,B int
}

//响应结构体
type ComputRespons struct {
	Pro  int   //乘积
	Que   int //商
	Rem  int //余数
}

//调用
func main()  {

	//链接远程的rpc
	conn,err:=jsonrpc.Dial("tcp","127.0.0.1:8081")
	if err!=nil{
		log.Fatal(err)
	}
	req:=ComputRequset{10,4}
	var res ComputRespons  //接收返回的参数
	err1:=conn.Call("Compute.Multip",req,&res)
	if err1!=nil{
		log.Fatal(err1)
	}
	fmt.Println("乘法==",res.Pro)

	//调用商
	err3:=conn.Call("Compute.Division",req,&res)
	if err3!=nil{
		log.Fatal(err3)
	}

	fmt.Printf("%d / %d = %d \n",req.A,req.B, res.Que)
	fmt.Printf("%d % %d = %d \n",req.A,req.B, res.Rem)
}

