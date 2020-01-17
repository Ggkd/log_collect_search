package main

import (
	"github.com/Ggkd/log_collect/etcd"
	"github.com/Ggkd/log_collect/kafka"
	"github.com/Ggkd/log_collect/taillog"
)

var (
	kafkaConfig		*kafka.Config
)

func main() {
	// 初始化etcd
	etcd.Init()

	// 加载kafka的配置
	kafka.LoadConfig()

	// 初始化kafka消费者
	kafka.InitConsumer()

	// 初始化kafka生产者
	kafka.InitProducer(kafkaConfig)

	// 从etcd获取日志收集的path 和 kafka的topic
	LogConf := etcd.GetConf(etcd.EtcdConfig.Key)

	// 开启任务
	taillog.Start(LogConf)



	ch := make(chan struct{}, 1)
	<- ch
}