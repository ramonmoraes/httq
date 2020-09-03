package infra

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gorilla/mux"
)

type KafkaHTTQ struct {
}

func (k *KafkaHTTQ) GetPrefix() string { return "kafka" }

var kafkaConfig *kafka.ConfigMap = &kafka.ConfigMap{
	"bootstrap.servers":            "localhost",
	"group.id":                     "myGroup",
	"auto.offset.reset":            "earliest",
	"queue.buffering.max.messages": 1,
}

func (k *KafkaHTTQ) GetMessage(w http.ResponseWriter, r *http.Request) {
	c, err := kafka.NewConsumer(kafkaConfig)
	if err != nil {
		invalidResponse(w, []byte(err.Error()))
		return
	}
	defer c.Close()

	params := mux.Vars(r)
	topic := params["key"]

	c.SubscribeTopics([]string{topic}, nil)
	fmt.Printf("Getting message from topic: [%s]\n", topic)
	content, err := c.ReadMessage(time.Second * 10)

	if err != nil {
		invalidResponse(w, []byte(err.Error()))
		return
	}

	validResponse(w, content.Value)
}

func (k *KafkaHTTQ) WriteMessage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	topic := params["key"]

	content, err := ioutil.ReadAll(r.Body)
	p, err := kafka.NewProducer(kafkaConfig)

	if err != nil {
		invalidResponse(w, []byte(err.Error()))
		return
	}
	defer p.Close()

	eventsChannel := make(chan kafka.Event)

	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          content,
	}, eventsChannel)

	event := <-eventsChannel
	validResponse(w, []byte(event.String()))
}
