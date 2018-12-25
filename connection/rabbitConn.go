package rabbitconn

import (
	"log"

	"github.com/streadway/amqp"
)

var rabbitConn *amqp.Connection
var rabbitChannel *amqp.Channel
var queueMap map[string]amqp.Queue

func RegisterQueue(queueName string) {
	if rabbitConn == nil {
		InitConn()
	}
	q, err := rabbitChannel.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	queueMap[queueName] = q
}

func Publish(queueName string, data []byte) {

}

func GetConsumer(queueName string) {

}

func GetConn() *amqp.Connection {
	if rabbitConn == nil {
		InitConn()
	}
	return rabbitConn
}

func InitConn() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	queueMap = make(map[string]amqp.Queue)
	rabbitConn = conn
	ch, _ := rabbitConn.Channel()
	rabbitChannel = ch
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
