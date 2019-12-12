/*
  Author ： CHR_崔贺然
  Time ： 2019.11.13
  Description ： go client连接kafka
*/
package kafka

// import (
// 	"fmt"
// 	"os"

// 	"github.com/Shopify/sarama"
// )

// func KafkaInit() {
// 	config := sarama.NewConfig()
// 	config.Producer.RequiredAcks = sarama.WaitForAll
// 	config.Producer.Partitioner = sarama.NewRandomPartitioner
// 	config.Producer.Return.Successes = true

// 	msg := &sarama.ProducerMessage{}
// 	msg.Topic = "nginx_log"
// 	msg.Value = sarama.StringEncoder("this is a good test")

// 	client, err := sarama.NewSyncProducer([]string{os.Getenv("KAFKA_CONF")}, config)
// 	if err != nil {
// 		fmt.Println("producer close, err:", err)
// 		return
// 	}
// 	defer client.Close()

// 	pid, offset, err := client.SendMessage(msg)
// 	if err != nil {
// 		fmt.Println("send message failed,", err)
// 		return
// 	}
// 	fmt.Println(pid, offset)
// }
