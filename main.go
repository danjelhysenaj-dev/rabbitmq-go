package main

import (
	"rabbitmq-go/receiver"
	"rabbitmq-go/sender"
)

func main() {
	// Calling Sender
	sender.Sender()
	// Calling Receiver
	receiver.Receiver()
}
