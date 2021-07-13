package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/fronge/ZFGFrameWork/logger"
	"gitlab.huayong.com/spider/grace/encoding"
)

var (
	Producer sarama.SyncProducer
)

func InitKafka() (err error) {
	// kafka参数设置
	conf := sarama.NewConfig()
	conf.Producer.Return.Successes = true                   // 同步模式
	conf.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	conf.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition

	// 连接kafka
	Producer, err = sarama.NewSyncProducer([]string{"192.168.0.101:9092", "192.168.0.101:9093", "192.168.0.101:9094"}, conf)
	if err != nil {
		logger.Info(err)
		return
	}
	logger.Infof("kafka connect success")

	return
}

func SendToKfK(topic string, value interface{}) (partition int32, offset int64, err error) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	prod, err := encoding.JSON.Marshal(value)
	if err != nil {
		return 0, 0, err
	}
	msg.Value = sarama.ByteEncoder(prod)
	fmt.Println(Producer)
	partition, offset, err = Producer.SendMessage(msg)
	return
}

func ReadFromKfk(topic string) {}
