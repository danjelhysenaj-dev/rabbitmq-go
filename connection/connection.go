package connection

import (
	"rabbitmq-go/utils"

	amqp "github.com/rabbitmq/amqp091-go"
)

// GetConnection return an AMQP connection
func GetConnection() *amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	return conn
}

// GetChannel is a function that returns a channel
func GetChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	return ch
}
