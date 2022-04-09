package util

import (
	go_errors "errors"
	"fmt"
	"sync"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

//阻止程序退出
var SubscribeWg sync.WaitGroup

//iot客户端
var IotClient MQTT.Client

//订阅消息线程
func SubscribeMessage(broker string, productId string, deviceId string, topic string, clientId string, userName string, password string) {
	SubscribeWg.Add(1)
	opts := MQTT.NewClientOptions().AddBroker(broker)
	opts.SetClientID(clientId)
	opts.SetUsername(userName)
	opts.SetPassword(password)
	opts.SetDefaultPublishHandler(handleMessage)
	opts.OnConnect = func(c MQTT.Client) {
		if token := c.Subscribe(fmt.Sprintf("%v/%v/%v", productId, deviceId, topic), 0, handleMessage); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	}

	//iot访问客户端
	IotClient = MQTT.NewClient(opts)
	if token := IotClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	//阻止线程退出
	SubscribeWg.Wait()
}

//发布数据到腾讯云iothub
func SendIoTHubData(productId string, deviceId string, topic string, payload string) error {
	if IotClient == nil {
		return go_errors.New("iothub客户端连接失败")
	}

	//发送topic
	sendTopic := fmt.Sprintf("%v/%v/%v", productId, deviceId, topic)
	token := IotClient.Publish(sendTopic, 0, false, payload)
	if token != nil {
		token.Wait()
	} else {
		return go_errors.New("iothub发送消息失败")
	}
	return nil
}

//订阅消息处理函数
var handleMessage MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {

}
