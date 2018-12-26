package main

import (
	"fmt"
	"log"

	rabbitmqConn "github.com/rabbitmqlab/connection"
)

func main() {
	fmt.Println("kerker")

	conn := rabbitmqConn.GetConn()
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	rabbitmqConn.RegisterQueue("hello")

	body := "Hello World!"
	err = rabbitmqConn.Publish("hello", []byte(body))
	failOnError(err, "Failed to publish a message")

	//Receive message
	msgs, err := rabbitmqConn.GetConsumer("hello")
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
