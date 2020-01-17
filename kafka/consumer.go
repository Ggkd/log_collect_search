package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
	"time"
)

var (
	TopicChan chan string		//topic通道
	Consumer sarama.Consumer	// 全局消费者
)

func Consume()  {
	for {
		select {
		case TopicChan := <- TopicChan:
			partitionList, err := Consumer.Partitions(TopicChan)		//根据topic找到所有的分区
			if err != nil {
				fmt.Println(err)
				return
			}
			//遍历所有的分区
			wg := sync.WaitGroup{}
			for _, partition := range partitionList {
				pc, err := Consumer.ConsumePartition(TopicChan, partition, sarama.OffsetNewest)
				if err != nil {
					fmt.Println(err)
				}
				defer pc.Close()
				//异步的消费数据
				wg.Add(1)
				go func(partitionConsumer sarama.PartitionConsumer) {
					for msg := range partitionConsumer.Messages(){
						fmt.Printf("topic:%s,  partiotion:%v, offset:%v, value:%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Value)
					}
					wg.Done()
				}(pc)
			}
			wg.Wait()
		default:
			time.Sleep(time.Microsecond)
		}
	}

}

func InitConsumer() {
	var err error
	config := sarama.NewConfig()
	address := KafkaConfig.Ip + ":" + KafkaConfig.Port
	Consumer, err = sarama.NewConsumer([]string{address}, config)
	if err != nil {
		fmt.Println("new consumer err: ", err)
		return
	}
	TopicChan = make(chan string, KafkaConfig.ChanSize)
	fmt.Println("===init kafka consumer success===")
	go Consume()
}

func SendToConsumeChan(topic string)  {
	TopicChan <- topic
}