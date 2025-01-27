# Introduction
This document will show how to setup locally and test RabbitMQ written in Golang

# Download locally RabbitMQ
```docker
docker run -it --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:4.0-management
```

# Run locally RabbitMQ

```docker
docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:4.0-management
```

it will be exposed at http://localhost:15672

# Build and Test
Note: Update line number 10 with guest and password guest in /connnection/connection.go file

```go
go run main.go
```

Run them seperately in different terminals where you first comment out the Receiver and run the Sender aftwerwards.

Below it shows that the message was consumed properly by the Receiver in the GUI of RabbitMQ.

