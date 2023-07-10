// This project is a simple package to help you to automate the creation of your message broker
// such as RabbitMQ, using the declarative way to create queues, exchanges and bindings.
package main

import (
	"github.com/nicollaspetrelli/declarator/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	amqpServerURL := "amqp://user:pass@localhost:5672"

	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}
	defer connectRabbitMQ.Close() // nolint:errcheck

	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer channelRabbitMQ.Close() // nolint:errcheck

	declarator := rabbitmq.NewDeclarator(channelRabbitMQ)

	declarator.DeclareFromFile("examples/hello-world-broker.json")
}
