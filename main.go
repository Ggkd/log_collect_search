package main

import (
	"github.com/Ggkd/log_collect/es"
	"github.com/Ggkd/log_collect/etcd"
	"github.com/Ggkd/log_collect/kafka"
	"github.com/Ggkd/log_collect/taillog"
)


func main() {
	// 初始化etcd
	etcd.Init()

	// 加载es的配置
	es.LoadEsConfig()

	// 初始化es
	es.Init()

	// 加载kafka的配置
	kafka.LoadConfig()

	// 初始化kafka消费者
	kafka.InitConsumer()

	// 初始化kafka生产者
	kafka.InitProducer()

	// 从etcd获取日志收集的path 和 kafka的topic
	LogConf := etcd.GetConf()

	// 开启任务
	taillog.Init(LogConf)

	// etcd Watch
	newConfChan := taillog.PushConfToChan()
	go etcd.WatchConf(newConfChan)



	ch := make(chan struct{}, 1)
	<- ch
}