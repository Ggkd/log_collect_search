package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

var ProductChan chan string

func SendToChan(msg string)  {
	ProductChan <- msg
}

func SendMsg(producer sarama.SyncProducer , topic string)  {
	defer producer.Close()
	//构造消息
	message := &sarama.ProducerMessage{}
	message.Topic = topic
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
	ProductChan = make(chan string, conf.MaxCap)
	fmt.Println("===init kafka producer success===")
	return producer
}