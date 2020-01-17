package main

import (
	"fmt"
	"github.com/Ggkd/log_collect/etcd"
	"github.com/Ggkd/log_collect/kafka"
	"github.com/Ggkd/log_collect/taillog"
	"github.com/Shopify/sarama"
	"github.com/hpcloud/tail"
)

var (
	t 				*tail.Tail
	producer		sarama.SyncProducer
	consumer		sarama.Consumer
	kafkaConfig		*kafka.Config
	//tailConfig		*taillog.Config
)

func main() {
	var err error
	// 加载etcd的配置
	err = etcd.LoadConfig()
	if err != nil {
		fmt.Println("load etcd config err :", err)
		return
	}
	fmt.Println("load etcd config  success")

	// 初始化etcd
	err = etcd.Init()
	if err != nil {
		fmt.Println("init etcd err :", err)
		return
	}
	fmt.Println("Init etcd success")

	// 开启etcd watch，获取日志配置
	go etcd.Watch()

	// 从监控的通道中取数据
	value := etcd.GetValue()

	// 加载tail的配置
	//tailConfig = taillog.LoadConfig()

	// 初始化tail
	go func() {
		t = taillog.InitTail(value)
	}()

	// 加载kafka的配置
	kafkaConfig = kafka.LoadConfig()

	// 初始化kafka生产者
	producer = kafka.InitProducer(kafkaConfig)

	// 初始化kafka消费者
	consumer = kafka.InitConsumer(kafkaConfig)

	// 消费者等待消费
	go kafka.Consume(consumer, value)

	// 等待接收通道消息并生产
	go kafka.SendMsg(producer, value)

	for line := range t.Lines {
		// 获取消息并发往通道
		go kafka.SendToChan(line.Text)
	}
	ch := make(chan struct{}, 1)
	<- ch
}