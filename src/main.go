package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)

// Prints a log message with error.
func printError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s, %s", msg, err)
	}
}

func main() {

	/* var virtualHost string = "test"
	   var username string = "test"
	   var password string = "test"
	   var exchangeName string = "test"
	   var host string = "test"
	   var queue string = "test"
	   var exchangeType = "test"
	   var routingKey = "test" */

	// To supply Vhost we need to pass instance of Config structure
	cfg := amqp.Config{Vhost: "TestVHost"}
	conn, err := amqp.DialConfig("amqp://rmquser:rmqpwd@localhost:5672/", cfg)
	printError(err, "Failed to connect to RabbitMQ")

	// Opens a unique server channel
	channel, err := conn.Channel()
	printError(err, "Failed to to open a channel")
	defer channel.Close()

	// Declare a queue to hold a message and deliver to consumers. It creates a queue if it doesn't already exist
	_, queueDeclarationErr := channel.QueueDeclare("TestQueueName", false, false, false, false, nil)
	printError(queueDeclarationErr, "Failed to declare Queue")

	// Bind and exchange to a queue
	channel.QueueBind("TestQueueName", "TestKeyName", "TestExchangeName", false, nil)

	// Prepare message
	msg := amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		Timestamp:    time.Now(),
		ContentType:  "text/plain",
		Body:         []byte("Hey this is from rmq_sender"),
	}

	// Publish message
	publishErr := channel.Publish("testExchangeName", "TestKeyName", false, false, msg)
	printError(publishErr, "Failed to publish message")

	defer conn.Close()
	fmt.Println("Message has sent succesfully.")
}
