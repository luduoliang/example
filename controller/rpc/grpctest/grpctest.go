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
