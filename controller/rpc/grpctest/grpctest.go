package grpctest

import (
	context "context"
	"fmt"
	"io"
	"net"
	"runtime"
	"time"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//测试服务
type TestFirstServerService struct {
	TestFirstServer
}

//开启grpc服务端
func GrpcServer(serverPort int) {
	//grpc服务
	grpcServer := grpc.NewServer(grpc.ConnectionTimeout((500/100)*time.Second + time.Millisecond*800))
	//注册服务
	RegisterTestFirstServer(grpcServer, &TestFirstServerService{})
	reflection.Register(grpcServer)
	//侦听端口
	address, err := net.Listen("tcp", fmt.Sprintf(":%v", serverPort))
	if err != nil {
		fmt.Println(fmt.Sprintf("grpc端口：%v开启失败：%v", serverPort, err.Error()))
		return
	}

	fmt.Println(fmt.Sprintf("grpc打开端口：%v", serverPort))

	//开启服务
	if err := grpcServer.Serve(address); err != nil {
		fmt.Println(fmt.Sprintf("grpc端口：%v开启失败%v", serverPort, err.Error()))
	}
}

//服务对应方法
func (s *TestFirstServerService) GetData(ctx context.Context, in *RequestFirst) (out *ResponseFirst, err error) {
	out = new(ResponseFirst)
	out.Message = "success"
	out.Code = "200"
	out.Data = []byte(in.Id + "ok")
	fmt.Println(out)
	return
}

//与客户端建立流通信方法
func (s *TestFirstServerService) Communite(stream TestFirst_CommuniteServer) error {
	var err error
	errChan := make(chan error, 1)
	messageReq := make(chan *RequestFirst, 1)
	go func(stream TestFirst_CommuniteServer) {
		defer close(errChan)
		defer close(messageReq)
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				errChan <- err
				runtime.Goexit()
			}
			if err != nil {
				errChan <- err
				runtime.Goexit()
			} else {
				messageReq <- in
			}
		}
	}(stream)

	for {
		select {
		case <-stream.Context().Done():
			goto exitlable
		case err = <-errChan:
			goto exitlable
		case in := <-messageReq:
			ProcessMessage(stream, in)
		}
	}
exitlable:
	return err
}

//处理客户端发来的消息
func ProcessMessage(stream TestFirst_CommuniteServer, in *RequestFirst) {
	stream.Send(&ResponseFirst{
		Code:    "200",
		Message: "success",
		Data:    []byte("end ok stream" + in.Id),
	})
}

//判断是否是无数据错误
func IsNoDataError(err error) bool {
	netErr, ok := err.(net.Error)
	if ok && netErr.Timeout() && netErr.Temporary() {
		return true
	}
	return false
}

/*
grpc客户端使用方法
main.go
func main() {
	go grpc_client.NewClient()
	go func() {
		for {
			fmt.Println(grpc_client.Client)
			if grpc_client.Stream != nil {
				err := grpc_client.Stream.Send(&pb.RequestFirst{
					Id: "9999",
				})
				if err != nil {
					fmt.Printf(err.Error())
				}

			}

			if grpc_client.Client != nil {
				resp, err := grpc_client.Client.GetData(context.Background(), &pb.RequestFirst{
					Id: "111",
				})
				if err != nil {
					fmt.Printf(err.Error())
				}
				fmt.Println("resp", string(resp.Data))
			}

			time.Sleep(2 * time.Second)
		}

	}()

	tiker := time.NewTicker(time.Minute * 30)
	for {
		<-tiker.C
	}
}


client.go：
var Conn *grpc.ClientConn
var Client pb.TestFirstClient
var Stream pb.TestFirst_CommuniteClient

func NewClient() {
	var err error
connect_next:
	fmt.Println("[grpc_client]开始初始化stream")
	for {
		Conn, err = grpc.Dial("127.0.0.1:7000", grpc.WithInsecure())
		if err != nil {
			fmt.Println("[grpc_client]创建grpc连接失败：", err)
			return
		}
		Client = pb.NewTestFirstClient(Conn)
		//ctx := context.Background()
		ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
		// 创建双向数据流
		Stream, err = Client.Communite(ctx)
		//创建流stream成功
		if err == nil {
			break
		}

		fmt.Println("[grpc_client]创建流stream失败：", err)
		if Stream != nil {
			Stream.CloseSend()
		}
		if Conn != nil {
			Conn.Close()
		}
		time.Sleep(time.Second * 2)
	}

	fmt.Println("[grpc_client]初始化stream成功")
	for {
		in, err := Stream.Recv()
		if err == io.EOF {
			continue
		}

		//断线重连，重新创建流stream
		if (err != nil) && (!isNoDataError(err)) {
			goto connect_next
		}
		//处理消息
		ProcessMessage(in)
	}

}

//处理服务端发来的消息
func ProcessMessage(in *pb.ResponseFirst) {
	fmt.Println(in)
}

//判断是否是无数据错误
func isNoDataError(err error) bool {
	netErr, ok := err.(net.Error)
	if ok && netErr.Timeout() && netErr.Temporary() {
		return true
	}
	return false
}
*/
