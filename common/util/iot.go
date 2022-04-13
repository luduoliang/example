package util

import (
	go_errors "errors"
	"fmt"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

//腾讯云iothub订阅消息，返回IOT客户端用于发布消息，传入回调函数用于接收订阅消息
func SubscribeIotMessage(broker string, productId string, deviceId string, topic string, clientId string, userName string, password string, procFunc func(client MQTT.Client, msg MQTT.Message)) (MQTT.Client, error) {
	opts := MQTT.NewClientOptions().AddBroker(broker)
	opts.SetClientID(clientId)
	opts.SetUsername(userName)
	opts.SetPassword(password)
	opts.SetDefaultPublishHandler(procFunc)
	opts.OnConnect = func(c MQTT.Client) {
		if token := c.Subscribe(fmt.Sprintf("%v/%v/%v", productId, deviceId, topic), 0, procFunc); token.Wait() && token.Error() != nil {
			return
		}
	}

	//iot访问客户端
	iotClient := MQTT.NewClient(opts)
	//连接IOT
	if token := iotClient.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
	return iotClient, nil
}

//腾讯云iothub发布消息
func PublishIotMessage(client MQTT.Client, productId string, deviceId string, topic string, payload interface{}) error {
	if client == nil {
		return go_errors.New("iothub客户端连接为nil")
	}
	//发送topic
	sendTopic := fmt.Sprintf("%v/%v/%v", productId, deviceId, topic)
	token := client.Publish(sendTopic, 0, false, payload)
	if token != nil {
		token.Wait()
	} else {
		return go_errors.New("iothub发送消息失败")
	}
	return nil
}
