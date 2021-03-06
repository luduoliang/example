package util

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	"github.com/segmentio/kafka-go"
)

//连接kafka服务器,返回kafka连接
func ConnectKafka(topic string, partition int, urls []string) (*kafka.Conn, error) {
	// KAFKA列表IP为空
	if len(urls) == 0 {
		err := errors.New("kafka url list count == 0")
		return nil, err
	}

	var kafkaList = make([]int, len(urls))
	var i int
	for i = 0; i < len(kafkaList); i++ {
		kafkaList[i] = i
	}
	// 随机算法，打乱Kafka连接序号
	kafkaList = ShuffleArrayInt(kafkaList)
	for i = 0; i < len(kafkaList); i++ {
		var url = urls[kafkaList[i]]
		url = strings.TrimSpace(url)
		if url == "" {
			continue
		}

		//与kafka服务器建立连接
		conn, err := kafka.DialLeader(context.Background(), "tcp", url, topic, partition)
		if err != nil {
			if conn != nil {
				conn.Close()
			}
			fmt.Println("send topic=>", topic, "(", url, ") err=>", err.Error())
			continue
		}
		//192.168.200.160 9092
		err = conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
		if err != nil {
			if conn != nil {
				conn.Close()
			}
			fmt.Println("SetWriteDeadline=>", topic, "(", url, ") err=>", err.Error())
			continue
		}
		return conn, nil
	}
	//返回发送成功
	return nil, errors.New("kafka初始化失败")
}

//发送kafka消息
//conn为kafka连接，message为发送的消息
func SendKafkaMessage(conn *kafka.Conn, message interface{}) error {
	// to produce messages
	var bs []byte
	bs, _ = json.Marshal(message)

	n, err := conn.WriteMessages(kafka.Message{Value: bs})
	if err != nil {
		return err
	}
	if n == 0 {
		return fmt.Errorf("write data len is:%v", n)
	}
	//返回发送成功
	return nil
}

//读取kafka消息，每次都读取该topic下全量数据，读取过的数据下次还会再读取
//conn为kafka连接，maxBytes一次最大读取字节数，procFunc为处理消息函数
func ReadKafkaMessage(conn *kafka.Conn, maxBytes int, procFunc func(interface{})) {
	// to produce messages
	for {
		//读取kafka消息，如果没有消息，则等待
		m, err := conn.ReadMessage(maxBytes)
		if err != nil {
			fmt.Println("==========kafka read=========>", err)
			time.Sleep(time.Second * 3)
			continue
		}
		//处理接收到的消息
		procFunc(m)
	}
}

//接收kafka消息，增量读取，读取过的消息就不再读取了
//urls为集群url地址列表，可以传单个
//不停的读取数据,读到数据后传给procFunc处理
func ReceiveKafkaMessage(topic string, groupId string, partition int, urls []string, procFunc func(interface{})) {
	if len(urls) == 0 {
		fmt.Println("----------->未配置kafka url")
		return
	}
	//初始化kafka
init_kafka:
	var r *kafka.Reader
	brokers := urls
	// make a new reader that consumes from topic-A
	//传Partition每次是全量读取
	/*r = kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		GroupID: "c1",
		//Partition:      partition,
		Topic:          topic,
		MinBytes:       10e3,        // 10KB
		MaxBytes:       10e6,        // 10MB
		CommitInterval: time.Second, // flushes commits to Kafka every second
	})*/

	r = kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,
		GroupID:        groupId,
		Partition:      partition,
		Topic:          topic,
		MinBytes:       10e3,        // 10KB
		MaxBytes:       10e6,        // 10MB
		CommitInterval: time.Second, // flushes commits to Kafka every second
	})
	defer r.Close()

	for {
		//读取kafka消息，如果没有消息，则等待
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("==========kafka Recv=========>", err)
			time.Sleep(time.Second * 3)
			goto init_kafka
		}
		//处理接收到的消息
		procFunc(m)
	}
}

//发送一次kafka消息，每发一条消息都会重新建立kafka连接，效率较底
//urls为集群url地址列表，可以传单个
func SendKafkaMessageOnce(topic string, partition int, urls []string, message interface{}) (err error) {
	// to produce messages
	var bs []byte
	// KAFKA列表IP为空
	if len(urls) == 0 {
		err = errors.New("kafka url list count == 0")
		return
	}

	var kafkaList = make([]int, len(urls))
	var i int
	for i = 0; i < len(kafkaList); i++ {
		kafkaList[i] = i
	}

	//发送kafka是否成功
	var isSendKafkaSuccess bool = false
	// 随机算法，打乱Kafka连接序号
	kafkaList = ShuffleArrayInt(kafkaList)
	for i = 0; i < len(kafkaList); i++ {
		var url = urls[kafkaList[i]]
		url = strings.TrimSpace(url)
		if url == "" {
			isSendKafkaSuccess = false
			continue
		}

		//与kafka服务器建立连接
		conn, err := kafka.DialLeader(context.Background(), "tcp", url, topic, partition)
		defer func() {
			if conn != nil {
				conn.Close()
			}
		}()
		if err != nil {
			isSendKafkaSuccess = false
			if conn != nil {
				conn.Close()
			}
			fmt.Println("send topic=>", topic, "(", url, ") err=>", err.Error())
			continue
		}
		//192.168.200.160 9092
		err = conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
		if err != nil {
			isSendKafkaSuccess = false
			if conn != nil {
				conn.Close()
			}
			fmt.Println("SetWriteDeadline=>", topic, "(", url, ") err=>", err.Error())
			continue
		}

		bs, _ = json.Marshal(message)

		len, err := conn.WriteMessages(kafka.Message{Value: bs})
		if err != nil {
			isSendKafkaSuccess = false
			if conn != nil {
				conn.Close()
			}
			fmt.Println("TX Length:", len, "err:", err)
			continue
		}

		if conn != nil {
			conn.Close()
		}

		isSendKafkaSuccess = true
		break
	}

	//如果发送kafka失败，返回错误
	if !isSendKafkaSuccess {
		if err == nil {
			err = errors.New("发送KAKFA失败")
		}
		return err
	}

	//返回发送成功
	return nil
}

//打乱数组元素
func ShuffleArrayInt(arr []int) (newArr []int) {
	// Fisher-Yates随机置乱算法
	rand.Seed(time.Now().Unix())
	var length int
	length = len(arr)
	if length <= 0 {
		return
	}
	newArr = make([]int, length)

	for i := len(arr) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		arr[i], arr[num] = arr[num], arr[i]
	}

	for i := 0; i < length; i++ {
		newArr[i] = arr[i]
	}
	return
}

//kafka消费者
func KafkaConsumer(topic string, partition int32, urls []string, procFunc func(*sarama.ConsumerMessage)) {
	//配置
	config := sarama.NewConfig()
	//接收失败通知
	config.Consumer.Return.Errors = true
	//设置使用的kafka版本,如果低于V0_10_0_0版本,消息中的timestrap没有作用.需要消费和生产同时配置
	config.Version = sarama.V0_11_0_0
	//新建一个消费者
	consumer, err := sarama.NewConsumer(urls, config)
	if err != nil {
		fmt.Println("error kafka get consumer")
		return
	}
	defer consumer.Close()

	//根据消费者获取指定的主题分区的消费者,Offset这里指定为获取最新的消息.
	partitionConsumer, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
	if err != nil {
		fmt.Println("error get partition consumer", err)
		return
	}
	defer partitionConsumer.Close()

	//循环等待接受消息.
	for {
		select {
		//接收消息通道和错误通道的内容.
		case msg := <-partitionConsumer.Messages():
			procFunc(msg)
		case err := <-partitionConsumer.Errors():
			fmt.Println(fmt.Sprintf("topic:%v,partition:%v,err:%v", err.Topic, err.Partition, err.Err.Error()))
		}
	}
}

//kafka生产者：异步消息模式，即发送的消息不需要被消费，还可以再发其它消息
func KafkaProducer(topic string, address []string, message interface{}) error {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second
	p, err := sarama.NewAsyncProducer(address, config)
	if err != nil {
		return err
	}
	defer p.Close()

	messageByte, _ := json.Marshal(message)
	msgProducer := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(messageByte),
	}

	p.Input() <- msgProducer
	return nil
}

//kafka生产者：同步消息模式，即发送的消息必须被消费了以后才可以再发其它消息
func KafkaSyncProducer(topic string, address []string, message interface{}) (int32, int64, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second
	p, err := sarama.NewSyncProducer(address, config)
	if err != nil {
		return 0, 0, err
	}
	defer p.Close()

	messageByte, _ := json.Marshal(message)
	msgProducer := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(messageByte),
	}

	part, offset, err := p.SendMessage(msgProducer)
	return part, offset, err
}
