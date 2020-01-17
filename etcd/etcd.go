package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

var Client *clientv3.Client	//Etcd全局客户端
var LogConfig []*LogEntry		// 所有的日志收集任务

type LogEntry struct {
	Path  string // tail日志的路径
	Topic string // kafka的topic
}

//初始化
func Init()  {
	// 加载etcd的配置
	err := LoadConfig()
	if err != nil {
		fmt.Println("load etcd config err :", err)
		return
	}
	fmt.Println("load etcd config  success")
	endpoints := EtcdConfig.Etcd.Ip + ":" + EtcdConfig.Etcd.Port
	Client, err = clientv3.New(clientv3.Config{
		Endpoints:[]string{endpoints},
		DialTimeout:time.Second*time.Duration(EtcdConfig.Etcd.DialTimeout),
	})
	if err != nil {
		fmt.Println("init etcd err :", err)
		return
	}
	LogConfig = make([] *LogEntry, EtcdConfig.TaskNum)
	fmt.Println("Init etcd success")
}

// 获取所有的日志配置
func GetConf(LogKey string) []*LogEntry {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := Client.Get(ctx, LogKey)
	if err != nil {
		fmt.Println("get err:", err)
	}
	for _, kv := range result.Kvs {
		logConfs := string(kv.Value)
		err := json.Unmarshal([]byte(logConfs), LogConfig)
		if err != nil {
			fmt.Println("unmarshal err : ", err)
			return nil
		}
	}
	return LogConfig
}