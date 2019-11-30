package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/bn_funds/utils"
	"github.com/google/logger"

	"github.com/streadway/amqp"
)

type Queue struct {
	Name    string
	Durable bool
	// Prefetch int
	Binding struct {
		Exchange string
		// Clean_start bool
		// Mannual_ack bool
	}
	Exchange struct {
		Name string
		Type string
	}
}

type AMQPQueue struct {
	Con   *amqp.Connection
	Ch    *amqp.Channel
	queue map[string]Queue
}

func (self *AMQPQueue) Init() {
	self.readJSON()
	self.connect()
	self.channel()
}

func (self *AMQPQueue) readJSON() {

	filename, _ := filepath.Abs("../jsons/amqp.json")
	byteValue, _ := ioutil.ReadFile(filename)

	json.Unmarshal(byteValue, &self.queue)
	// log.Printf("JSONData: %v\n", self.queue)
}

func (self *AMQPQueue) connect() {
	connection, err := amqp.Dial(utils.GetEnv("AMQP_URL", ""))
	if err != nil {
		logger.Error("Failed to connect to RabbitMQ: %+v", err)
	}
	// defer connection.Close()
	self.Con = connection

}

func (self *AMQPQueue) channel() {

	channel, err := self.Con.Channel()
	failOnError(err, "Failed to open a channel")
	// defer channel.Close()
	self.Ch = channel
}

func (self *AMQPQueue) Subscribe(
	id string) <-chan amqp.Delivery {
	if self.queue[id].Binding.Exchange != "" && self.queue[id].Exchange.Name != "" {
		err := self.Ch.ExchangeDeclare(
			self.queue[id].Exchange.Name, // name
			self.queue[id].Exchange.Type, // type
			true,  // durable
			false, // auto-deleted
			false, // internal
			false, // no-wait
			nil,   // arguments
		)
		failOnError(err, "Failed to declare an exchange")
	}
	q, err := self.Ch.QueueDeclare(
		self.queue[id].Name,    // name
		self.queue[id].Durable, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	if self.queue[id].Binding.Exchange != "" && self.queue[id].Exchange.Name != "" {
		err := self.Ch.QueueBind(
			self.queue[id].Name, // queue name
			"",                  // routing key
			self.queue[id].Exchange.Name, // exchange
			false,
			nil)
		failOnError(err, "Failed to bind a queue")
	}

	msgs, err := self.Ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	failOnError(err, "Failed to register a consumer")
	return msgs
}

func (self *AMQPQueue) Enqueue(id string, payload []byte, attrs map[string]interface{}) {
	var persistent int
	if attrs != nil {
		if attrs["persistent"].(bool) == true {
			persistent = 2 //persistent
		} else {
			persistent = 1 //non-persistent
		}
	}
	err := self.Ch.Publish(
		self.queue[id].Exchange.Name, // exchange
		self.queue[id].Name,          // routing key
		false,                        // mandatory
		false,                        // immediate
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         []byte(payload),
			DeliveryMode: uint8(persistent),
		})
	// log.Printf(" [x] Sent %s", payload)
	failOnError(err, "Failed to publish a message")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
