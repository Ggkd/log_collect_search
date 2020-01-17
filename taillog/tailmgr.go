package taillog

import "github.com/Ggkd/log_collect/etcd"

type TaskMgr struct {
	logEntry []*etcd.LogEntry
}

// 遍历每一个任务
func Start(LogConf []*etcd.LogEntry)  {
	// 遍历获取每一个task
	for _, logConfig := range LogConf {
		// 初始化tail
		NewLogTask(logConfig.Path, logConfig.Topic)
	}
}