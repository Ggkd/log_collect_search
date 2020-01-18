package taillog

import (
	"context"
	"fmt"
	"github.com/Ggkd/log_collect/kafka"
	"github.com/hpcloud/tail"
)

type LogTask struct {
	path 			string		// 日志路径
	topic 			string		// kafka主题
	instance 		*tail.Tail	// Tail实例
	ctx 			context.Context			// 用于停止task
	cancelFunc		context.CancelFunc		// 用于停止task
}

// 构造函数
func NewLogTask(path, topic string) *LogTask {
	logTask := new(LogTask)
	ctx, cancel := context.WithCancel(context.Background())
	logTask.path = path
	logTask.topic = topic
	logTask.ctx = ctx
	logTask.cancelFunc = cancel
	logTask.InitTail(path)
	return logTask
}

// tail初始化
func (lt *LogTask)InitTail(path string) {
	t, err := tail.TailFile(path, tail.Config{Follow: true})
	lt.instance = t
	if err != nil {
		fmt.Println("tail file err: ", err)
		return
	}
	fmt.Println("====init tail config success====")
	//后台读取log
	go lt.Run()
}

// 想kafka发送数据
func (lt *LogTask) Run()  {
	for {
		select {
		// 等待退出信号
		case <- lt.ctx.Done():
			fmt.Println("配置------------->"+lt.path+"-----"+lt.topic+" 已退出")
			return
		case line := <- lt.instance.Lines:
			// 发送数据到kafka
			kafka.SendToProducerChan(lt.topic, line.Text)
			kafka.SendToConsumeChan(lt.topic)
		}
	}
}