package rpctest

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

//基于http2服务
type RpcHttpService struct {
}

//基于tcp服务
type RpcTcpService struct {
}

//传输方法
type TestData struct {
	Id   string
	Name string
}

//服务RpcHttpService的方法
func (obj *RpcHttpService) TestProc(args *TestData, reply *string) error {
	*reply = args.Id + "_" + args.Name
	return nil
}

//服务RpcTcpService的方法
func (obj *RpcTcpService) TestProc(args *TestData, reply *string) error {
	*reply = args.Id + "_" + args.Name
	return nil
}

//开启基于http2服务的rpc服务端
func RpcHttpServer(serverPort int) {
	err := rpc.RegisterName("RpcHttpService", new(RpcHttpService))
	if err != nil {
		fmt.Println(fmt.Sprintf("rpc端口：%v开启失败：%v", serverPort, err.Error()))
		return
	}
	rpc.HandleHTTP()

	fmt.Println(fmt.Sprintf("rpc打开端口：%v", serverPort))
	if err = http.ListenAndServe(fmt.Sprintf(":%v", serverPort), nil); err != nil {
		fmt.Println(fmt.Sprintf("rpc端口：%v，服务开启失败", serverPort))
	}
}

//开启基于tcp服务的rpc服务端
func RpcTcpServer(serverPort int) {
	err := rpc.RegisterName("RpcTcpService", new(RpcTcpService))
	if err != nil {
		fmt.Println(fmt.Sprintf("rpc端口：%v开启失败：%v", serverPort, err.Error()))
		return
	}

	listen, err := net.Listen("tcp", fmt.Sprintf(":%v", serverPort))
	if err != nil {
		fmt.Println(fmt.Sprintf("rpc端口：%v开启失败：%v", serverPort, err.Error()))
		return
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		defer conn.Close()

		if err != nil {
			fmt.Println(fmt.Sprintf("rpc端口：%v接收数据失败：%v", serverPort, err.Error()))
			return
		}
		rpc.ServeConn(conn)
	}
}

/*
备注：
基于tcp客户端调用方法
//测试传输方法
type TestData struct {
	Id   string
	Name string
}

func main() {
	client, err := rpc.Dial("tcp", ":7777")
	if err != nil {
		log.Fatal("dialing", err)
	}

	var reply string

	err = client.Call("RpcTcpService.TestProc", TestData{Id: "2", Name: "你爷爷"}, &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}

基于http客户端调用方法

//测试传输方法
type TestData struct {
	Id   string
	Name string
}

func main() {
	client, err := rpc.DialHTTP("tcp", ":6666")
	if err != nil {
		log.Fatal("client error:", err)
	}
	args := TestData{Id: "1", Name: "zhangshan哈哈"}
	var reply string
	err = client.Call("RpcHttpService.TestProc", args, &reply)
	if err != nil {
		log.Fatal("arithmetic error:", err)
	}
	fmt.Println(reply)
}
*/
