package etcd

import (
	"github.com/coreos/etcd/clientv3"
	"time"
)

var Client *clientv3.Client	//全局客户端
var ValueChan chan string	//存放数据的通道

// 定义etcd put的数据； 日志的路径和kafka的topic
type PutValue struct {
	Path 	string		`json:"path"`
	Topic	string		`json:"topic"`
}

//初始化
func Init() error {
	var err error
	endpoints := EtcdConfig.Etcd.Ip + ":" + EtcdConfig.Etcd.Port
	Client, err = clientv3.New(clientv3.Config{
		Endpoints:[]string{endpoints},
		DialTimeout:time.Second*time.Duration(EtcdConfig.Etcd.DialTimeout),
	})
	// 初始化一个接收key的value通道
	ValueChan = make(chan string, EtcdConfig.Etcd.ChanSize)
	return err
}