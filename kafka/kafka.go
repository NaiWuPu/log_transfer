package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log_transfer/es"
)

// 初始化kafka消费者，往ES发送数据



// Init 初始化 client
func Init(address []string, topic string) (err error) {
	consumer, err := sarama.NewConsumer(address, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return err
	}
	// 根据topic获取所有的分区
	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		fmt.Printf("failed to get list of partition:err%v,\n", err)
		return err
	}
	fmt.Println("分区列表", partitionList)
	for partition := range partitionList {
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition :%d,err:%v\n", partition, err)
			return err
		}
		defer pc.AsyncClose()
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v value:%v\n", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
				// 发送给ES
				var ld = es.LogData{
					Topic: topic,
					Data: string(msg.Value),
				}
				es.Send2EsChan(&ld)
			}
		}(pc)
	}
	return err
}
