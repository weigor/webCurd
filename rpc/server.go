package main



import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)
//计算结构体
type Compute struct {
}

//接收参数
type ComputRequset struct {
	A,B int
}

//响应结构体
type ComputRespons struct {
	Pro  int   //乘积
	Que   int //商
	Rem  int //余数
}

//计算方法
func (this *Compute)Multip(req ComputRequset,res *ComputRespons)error  {
	res.Pro=req.A*req.B
	return nil
}

//商和余数

func (this *Compute)Division(req ComputRequset,res *ComputRespons)error  {
	if req.B==0{
		return  errors.New("除数不能为0")
	}
	res.Que=req.A/req.B
	res.Rem=req.A%req.B
	return nil
}

func main()  {

	//注册服务
	err:=rpc.Register(new(Compute))
	if err!=nil{
		log.Fatal(err)
	}

	//监听服务
	lis,err1:=net.Listen("tcp","127.0.0.1:8081")
	if err1!=nil{
		log.Fatal(err1)
	}

	for  {
		conn,err2:=lis.Accept()
		if err2!=nil{
			continue
		}
		// 携程
		go func(conn net.Conn) {
			fmt.Println("New a server")
			jsonrpc.ServeConn(conn)
		}(conn)
	}




}

