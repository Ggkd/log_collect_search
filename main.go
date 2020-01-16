package main

import (
	"github.com/Ggkd/log_collect/kafka"
	"github.com/Ggkd/log_collect/taillog"
)

func main() {
	// 加载tail的配置
	tailConfig := taillog.LoadConfig()
	// 初始化tail
	t := taillog.InitTail(tailConfig)
	// 加载kafka的配置
	kafkaConfig := kafka.LoadConfig()
	// 初始化kafka生产者
	producer := kafka.InitProducer(kafkaConfig)
	// 初始化kafka消费者
	consumer := kafka.InitConsumer(kafkaConfig)
	// 消费者等待消费
	go kafka.Consume(consumer, kafkaConfig.Topic)
	// 发送消息
	go kafka.SendMsg(producer, kafkaConfig.Topic)
	for line := range t.Lines {
		// 获取消息并发往通道
		go kafka.SendToChan(line.Text)
	}
	ch := make(chan struct{}, 1)
	<- ch
}