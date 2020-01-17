package etcd

import (
	"context"
	"fmt"
	"time"
)

// 监控数据的变化，并将数据发往通道
func Watch() {
	watchChan := Client.Watch(context.Background(), EtcdConfig.Key)
	for {
		for event := range watchChan {
			for _, ev := range event.Events {
				fmt.Printf("type:%s,  key:%s, value:%s \n", ev.Type, ev.Kv.Key, ev.Kv.Value)
				ValueChan <- string(ev.Kv.Value)
			}
		}
	}
}


// 从通道取数据
func GetValue() string {
	for {
		select {
		case value := <- ValueChan:
			return value
		default:
			time.Sleep(time.Microsecond)
		}
	}
}