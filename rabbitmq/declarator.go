package rabbitmq

import (
	"encoding/json"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
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

	var definition BrokerDefinition

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &definition)

	self.DeclareExchanges(definition.Exchanges)
	self.DeclareQueues(definition.Queues)
	self.DeclareBindings(definition.Bindings)

	log.Info("[RabbitMQ] [Declarator] All exchanges, queues and bindings declared from file " + filePath + " successfully!")

	defer jsonFile.Close()
}
