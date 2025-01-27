package sender

import (
	"context"
	"log"
	"rabbitmq-go/connection"
	"rabbitmq-go/utils"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Sender is a function that sends messages to the receiver
func Sender() {
	// Establish connection to RabbitMQ
	conn := connection.GetConnection()
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {
			log.Printf("Failed to close connection: %s", err)
		}
	}(conn)

	// Opening a channel for communication between Publisher and Consumer
	ch := connection.GetChannel(conn)
	defer func(conn *amqp.Channel) {
		err := ch.Close()
		if err != nil {
			log.Printf("Failed to close connection: %s", err)
		}
	}(ch)

	q, err := ch.QueueDeclare(
		"hello", //name
		false,   //durable
		false,   //delete when unused
		false,   //exclusive
		false,   //no-wait
		nil,     //args
	)
	utils.FailOnError(err, "Failed to declare queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := "Hello World"
	err = ch.PublishWithContext(ctx, "", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})
	utils.FailOnError(err, "Failed to Publish message")
	log.Printf("[x] Sent %s\n", body)

}
