// Package rabbitmq provides a declarator for rabbitmq message broker
package rabbitmq

import (
	"encoding/json"
	"io"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
)

type Declarator struct {
	conn *amqp.Channel
}

// NewDeclarator creates a new declarator
func NewDeclarator(conn *amqp.Channel) *Declarator {
	return &Declarator{
		conn: conn,
	}
}

// DeclareExchanges declares all exchanges ,queues and bindings from a broker definition json file
func (self *Declarator) DeclareFromFile(filePath string) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		log.Error("[RabbitMQ] [Declarator] Error when trying to open file " + filePath)
		return
	}

	defer jsonFile.Close() // nolint:errcheck

	var definition BrokerDefinition

	byteValue, _ := io.ReadAll(jsonFile)
	_ = json.Unmarshal(byteValue, &definition)

	self.DeclareExchanges(definition.Exchanges)
	self.DeclareQueues(definition.Queues)
	self.DeclareBindings(definition.Bindings)

	log.Info("[RabbitMQ] [Declarator] All exchanges, queues and bindings declared from file " + filePath + " successfully!")
}
