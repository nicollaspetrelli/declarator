# GoLang RabbitMQ Declarator Package

[![Go Reference](https://pkg.go.dev/badge/github.com/nicollaspetrelli/declarator.svg)](https://pkg.go.dev/github.com/nicollaspetrelli/declarator) [![Go Report Card](https://goreportcard.com/badge/github.com/nicollaspetrelli/declarator)](https://goreportcard.com/report/github.com/nicollaspetrelli/declarator) [![Go Version badge](https://img.shields.io/badge/Go-1.20.5-blue.svg)]()

# The Project

This project is a simple package to help you to automate the creation of your message broker such as RabbitMQ, using the declarative way to create queues, exchanges and bindings.

Initially is only supported RabbitMQ, but I'm planning to add support to other message brokers in the future.

# Why?

I created this package to help me to automate the creation of my message broker, because I have a lot of microservices and I need to create a lot of queues, exchanges and bindings. So, I created this package to help me to automate this process.

# How it works?

Actually, this package has two ways to declare your message broker.

### Using a file

This package will read a json file with the declarations and will create the queues, exchanges and bindings.

The format of declaration file is the same as RabbitMQ export file.

> Has a example of declaration file is in the examples folder, If you want more information about the format of declaration file can be found in the [RabbitMQ Documentation](https://www.rabbitmq.com/management-cli.html#export-import)

_Example_

```go
package main

import (
    "fmt"
    "log"

    "github.com/nicollaspetrelli/declarator/rabbitmq"
)

func main() {
    var rabbitConnection *amqp.Channel

    // Create a new declarator passing the connection
    declarator := rabbitmq.NewDeclarator(rabbitConnection)

    // Use declarator to declare from a json definitions file
    declarator.DeclareFromFile("examples/hello-world-broker.json")
}
```

### Using separated functions

Also you can use the separated functions to declare in your code.

_Example_

```go
package main

import (
    "fmt"
    "log"

    "github.com/nicollaspetrelli/declarator/rabbitmq"
)

func main() {
    var rabbitConnection *amqp.Channel

    // Create a new declarator passing the connection
    declarator := rabbitmq.NewDeclarator(rabbitConnection)

    // Declare a queue
    declarator.DeclareQueue(rabbitmq.Queue{
        Name: "hello-world-queue",
        Durable: true,
        AutoDelete: false,
        Exclusive: false,
        NoWait: false,
        Args: nil,
    })

    // Declare an exchange
    declarator.DeclareExchange(rabbitmq.Exchange{
        Name: "hello-world-exchange",
        Type: "direct",
        Durable: true,
        AutoDelete: false,
        Internal: false,
        NoWait: false,
        Args: nil,
    })

    // Declare a binding
    declarator.DeclareBinding(rabbitmq.Binding{
        Queue: "hello-world-queue",
        Exchange: "hello-world-exchange",
        RoutingKey: "hello-world-routing-key",
        NoWait: false,
        Args: nil,
    })
}
```

# Installing

How to install the package?

> `go get github.com/nicollaspetrelli/declarator`

Now you can use the package in your project following the examples above or in the main.go file in root of the project.

> Also you can use the docker-compose file to run a RabbitMQ instance to test the package.

# Future Work

- [ ] Make RabbitMQ Connection optional, passing DSN and create a new connection
- [ ] Add unit tests
- [ ] Add integration tests

# Contributing

Want to contribute? Great!

Just follow the steps below:

- Fork the project
- Create a branch with your feature
- Commit your changes
- Push your branch
- Create a new Pull Request

## Development

### Requirements

- Install [Golang](https://golang.org)
- Install [GolangCI-Lint](https://golangci-lint.run/)
- Install [docker](https://docs.docker.com/install/)
- Install [docker-compose](https://docs.docker.com/compose/install/)

### Makefile

Please run the make target below to see the provided targets.

```sh
$ make help
```

# License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Authors

- **Nicollas Petrelli** - _Initial work_ - [Contact me](mailto:me@nicollas.dev)
