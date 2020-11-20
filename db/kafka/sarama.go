package kafka

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Shopify/sarama"
)

const (
	Topic = "topic_name"
)

var (
	addresses = []string{"manage28.domain.com:30091"}
)

// Main kafka consumer
func Main() {
	consumer, err := sarama.NewConsumer(addresses, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	partitionList, err := consumer.Partitions(Topic) // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	fmt.Println(partitionList)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition(Topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		// defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%s Value:%s", msg.Partition, msg.Offset, msg.Key, msg.Value)
			}
		}(pc)
	}

	c := make(chan os.Signal)
	// 监听指定信号
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGUSR1, syscall.SIGUSR2)
	fmt.Println("启动了程序")
	s := <-c
	fmt.Println("收到信号:", s)
}
