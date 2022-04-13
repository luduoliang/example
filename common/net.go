package common

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"errors"
)

//POST请求
func HttpPost(url string, jsonParam string) (res []byte, err error) {
	var timeout time.Duration = time.Second
	var tempClinet = http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: timeout,
			}).DialContext,
		},
	}
	reqest, err := http.NewRequest("POST", url, bytes.NewReader([]byte(jsonParam)))
	if err != nil {
		return
	}
	reqest.Header.Set("Content-Type", "application/json; encoding=utf-8")
	reqest.Header.Set("Connection", "close")
	resp, err := tempClinet.Do(reqest)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return
	}
	var body []byte
	if resp.StatusCode == 200 {
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}
	} else {
		err = fmt.Errorf("StatusCode=%v error", resp.StatusCode)
	}
	res = body
	return
}

//GET请求
func HttpGet(url string, params map[string]string, headers map[string]string) ([]byte, error) {
	c := http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: time.Second,
			}).DialContext,
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()
	//new request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []byte{}, errors.New("new request is fail ")
	}
	req.Header.Set("Content-Type", "application/json; encoding=utf-8")
	req.Header.Set("Connection", "close")
	//add params
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	//add headers
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}
	//http client
	fmt.Printf("Go %s URL : %s \n", http.MethodGet, req.URL.String())
	do, err := c.Do(req.WithContext(ctx))
	if err != nil {
		return []byte{}, errors.New("do req failed.")
	}
	defer do.Body.Close()

	bytesBody, err := ioutil.ReadAll(do.Body)
	if err != nil {
		return []byte{}, errors.New("read body failed")
	}

	fmt.Printf("HttpGet receive : %s \n", string(bytesBody))
	return bytesBody, nil
}

//建立TCP服务端，并接收处理返回
//serverPort为服务端端口，buffSize为一次读取大小字节数，procFunc为接收到消息处理
func TcpServer(serverPort int, buffSize int, procFunc func([]byte, net.Conn)) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", serverPort))
	if err != nil {
		fmt.Println("TCP服务端监听端口失败：", err.Error())
		return
	}
	defer listener.Close()
	fmt.Println(fmt.Sprintf("TCP打开端口：%v", serverPort))

	for {
		//循环接入所有客户端得到专线连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("客户端连接失败：", err.Error())
		}
		defer conn.Close()

		fmt.Println(fmt.Sprintf("客户端：%v，与%v端口建立连接", conn.RemoteAddr(), serverPort))

		//一个客户端开启一个线程处理
		go func() {
			var chanExit chan int = make(chan int, 1)

			go func() {
				//创建消息缓冲区
				buffer := make([]byte, buffSize)
				for {
					///---一个完整的消息回合
					//读取客户端发来的消息放入缓冲区
					n, err := conn.Read(buffer)
					//如果读取该客户端消息失败，认为它断线了
					if err != nil {
						fmt.Println(fmt.Sprintf("客户端：%v，读取消息失败：", conn.RemoteAddr(), err.Error()))
						chanExit <- 1
					}

					//处理客户端消息
					procFunc(buffer[0:n], conn)
				}
			}()

			select {
			case <-chanExit:
				fmt.Println(fmt.Sprintf("客户端：%v，断开连接", conn.RemoteAddr()))
				break
			}
		}()

	}
}

//建立UDP服务端，并接收处理返回
//serverPort为服务端端口，buffSize为一次读取大小字节数，procFunc为接收到消息处理
func UdpServer(serverPort int, buffSize int, procFunc func([]byte, *net.UDPConn, *net.UDPAddr)) {
	udpAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%v", serverPort))
	if err != nil {
		fmt.Println("udp服务端监听端口失败：", err.Error())
		return
	}
	fmt.Println(fmt.Sprintf("UDP打开端口：%v", serverPort))

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("udp服务端监听端口失败：", err.Error())
		return
	}
	defer conn.Close()

	for {
		//循环读取UDP数据
		var buffer = make([]byte, buffSize)
		//接收客户端发送的数据
		n, clientaAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("udp服务端读取数据失败：", err.Error())
			continue
		}
		if n <= 0 {
			fmt.Println("udp服务端读取数据为空")
			continue
		}

		//处理客户端消息
		procFunc(buffer[0:n], conn, clientaAddr)
	}
}

/*
tcp客户端建立长连接方法

func main() {
	//断线重连
relink:
	//拨号远程地址，简历tcp连接
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println(err)
		if conn != nil {
			conn.Close()
		}
		time.Sleep(2 * time.Second)
		goto relink
	}

	//预先准备消息缓冲区
	buffer := make([]byte, 1024)
	//准备命令行标准输入
	reader := bufio.NewReader(os.Stdin)

	//断线控制
	var connChan chan int = make(chan int, 1)

	//写消息线程
	go func() {
		for {
			lineBytes, _, _ := reader.ReadLine()
			_, err := conn.Write(lineBytes)
			if err != nil {
				connChan <- 1
				break
			}
		}
	}()

	//读消息线程
	go func() {
		for {
			n, err := conn.Read(buffer)
			if err != nil {
				connChan <- 1
				break
			}
			serverMsg := string(buffer[0:n])
			fmt.Println("服务端msg", serverMsg)
			if serverMsg == "bye" {
				break
			}
		}
	}()

	select {
	case <-connChan:
		goto relink
	}
}

UDP客户端建立方法
func main() {
	//udp服务端口号
	udpAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%v", 7777))
	if err != nil {
		fmt.Println(err)
		return
	}
	udpClient, err := net.DialUDP("udp4", nil, udpAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	if udpClient != nil {
		defer udpClient.Close()
	}

	//发消息
	go func() {
		for {
			var msg MessageData
			msg.MessageType = "1"
			msg.MessageData = "1111111111"
			var reqJSON, _ = json.Marshal(msg)

			// 发送到服务端
			_, err = udpClient.Write(reqJSON)
			if err != nil {
				fmt.Println("UDP=>", err, string(reqJSON))
				break
			}
			time.Sleep(time.Second * 2)
		}
	}()

	//收消息
	go func() {
		for {
			var data = make([]byte, 4096)
			//接收客户端发送的数据
			n, _, err := udpClient.ReadFromUDP(data)
			if err != nil {
				fmt.Println("failed read udp msg, error: " + err.Error())
				continue
			}
			if n <= 0 {
				fmt.Println("<0")
				continue
			}
			fmt.Println("udp接收=>", n)

			//解析数据
			var msgData MessageData
			defer func() {
				err := recover()
				if err != nil {
				}
			}()
			err = json.Unmarshal(data[:n], &msgData)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("[收到的消息为]", msgData)

			time.Sleep(time.Second * 2)
		}

	}()

	tiker := time.NewTicker(time.Minute * 30)
	for {
		<-tiker.C
	}
}

//消息格式
type MessageData struct {
	MessageType string `json:"message_type"`
	MessageData string `json:"message_data"`
}

*/
