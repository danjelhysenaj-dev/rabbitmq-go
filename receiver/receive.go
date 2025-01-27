package receiver

import (
	"log"
	"rabbitmq-go/connection"
	"rabbitmq-go/utils"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Receiver() {

	// Establish connection to RabbitMQ
	conn := connection.GetConnection()
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
	}(conn)

	// Opening a channel for communication between Publisher and Consumer
	ch := connection.GetChannel(conn)
	defer func(conn *amqp.Channel) {
		err := ch.Close()
		if err != nil {
			log.Println(err)
		}
	}(ch)

	// Declare a queue
	q, err := ch.QueueDeclare(
		"hello", //name of the queue
		false,   // durable
		false,   // delete when-used
		false,   //exclusive
		false,   //no-wait
		nil,     //args
	)

	utils.FailOnError(err, "Failed to create a queue")

	// registering the consumer
	msg, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	utils.FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msg {
			log.Printf("received a message: %s", d.Body)
		}
	}()
	log.Printf(" [*] Waiting for messages, To exit press CTRL+C")
	<-forever
}
