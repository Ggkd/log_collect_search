package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/Ggkd/log_collect/etcd"
	"github.com/Shopify/sarama"
	"time"
)

var ProductChan chan string

// 从taillog获取到的日志 发往消息通道
func SendToChan(msg string)  {
	ProductChan <- msg
}

// 从通道获取消息并生产
func SendMsg(producer sarama.SyncProducer , value string)  {
	defer producer.Close()
	var putValue = new(etcd.PutValue)
	json.Unmarshal([]byte(value), putValue)
	//构造消息
	message := &sarama.ProducerMessage{}
	message.Topic = putValue.Topic
	for {
		select {
		case msg := <- ProductChan:
			message.Value = sarama.StringEncoder(msg)
			_, _, err := producer.SendMessage(message)
			if err != nil {
				fmt.Println(err)
				return
			}
			//fmt.Printf("partition : %v,  offset : %v\n", partition, offset)
		default:
			time.Sleep(time.Second)
		}
	}
}

//初始化生产者
func InitProducer(conf *Config) sarama.SyncProducer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	//连接kafka
	address := conf.Ip + ":" + conf.Port
	producer, err := sarama.NewSyncProducer([]string{address}, config)
	if err != nil {
		fmt.Println("init producer err : ", err)
		return nil
	}
	// 初始化ConsumeChan
	ProductChan = make(chan string, conf.ChanSize)
	fmt.Println("===init kafka producer success===")
	return producer
}