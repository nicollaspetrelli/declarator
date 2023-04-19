package rabbitmq

import (
	"github.com/streadway/amqp"
)

// BrokerDefinition is the structure of a broker definition json file
type BrokerDefinition struct {
	Queues    []Queue    `json:"queues"`
	Exchanges []Exchange `json:"exchanges"`
	Bindings  []Binding  `json:"bindings"`
}

// ExchangeArguments is the structure of the arguments of an exchange
type ExchangeArguments struct {
	AlternateExchange string `json:"alternate-exchange"`
}

// Exchange is the structure of an exchange
type Exchange struct {
	Name       string            `json:"name"`
	Vhost      string            `json:"vhost"`
	Type       string            `json:"type"`
	Durable    bool              `json:"durable"`
	AutoDelete bool              `json:"auto_delete"`
	Internal   bool              `json:"internal"`
	NoWait     bool              `json:"no_wait"`
	Arguments  ExchangeArguments `json:"arguments"`
}

// QueueArguments is the structure of the arguments of a queue
type QueueArguments struct {
	XMessageTTL           int32  `json:"x-message-ttl"`
	XExpires              int32  `json:"x-expires"`
	XOverflow             string `json:"x-overflow"`
	XSingleActiveConsumer bool   `json:"x-single-active-consumer"`
	XDeadLetterExchange   string `json:"x-dead-letter-exchange"`
	XDeadLetterRoutingKey string `json:"x-dead-letter-routing-key"`
	XMaxLength            int32  `json:"x-max-length"`
	XMaxLengthBytes       int32  `json:"x-max-length-bytes"`
	XMaxPriority          int32  `json:"x-max-priority"`
	XQueueMode            string `json:"x-queue-mode"`
	XQueueMasterLocator   string `json:"x-queue-master-locator"`
}

// Queue is the structure of a queue
type Queue struct {
	Name       string         `json:"name"`
	Vhost      string         `json:"vhost"`
	Durable    bool           `json:"durable"`
	AutoDelete bool           `json:"auto_delete"`
	Exclusive  bool           `json:"exclusive"`
	NoWait     bool           `json:"no_wait"`
	Arguments  QueueArguments `json:"arguments"`
}

// BindArguments is the structure of the arguments of a binding
type BindArguments struct{}

// Binding is the structure of a binding
type Binding struct {
	Source          string        `json:"source"`
	Vhost           string        `json:"vhost"`
	Destination     string        `json:"destination"`
	DestinationType string        `json:"destination_type"`
	RoutingKey      string        `json:"routing_key"`
	NoWait          bool          `json:"no_wait"`
	Arguments       BindArguments `json:"arguments"`
}

// GetArguments returns the arguments of an exchange
func (self *ExchangeArguments) GetArguments() amqp.Table {
	arguments := amqp.Table{}

	addNonEmpty(arguments, "alternate-exchange", self.AlternateExchange, "string")

	return arguments
}

// GetArguments returns the arguments of a queue
func (self *QueueArguments) GetArguments() amqp.Table {
	arguments := amqp.Table{}

	addNonEmpty(arguments, "x-message-ttl", self.XMessageTTL, "int32")
	addNonEmpty(arguments, "x-expires", self.XExpires, "int32")
	addNonEmpty(arguments, "x-overflow", self.XOverflow, "string")
	addNonEmpty(arguments, "x-single-active-consumer", self.XSingleActiveConsumer, "bool")
	addNonEmpty(arguments, "x-dead-letter-exchange", self.XDeadLetterExchange, "string")
	addNonEmpty(arguments, "x-dead-letter-routing-key", self.XDeadLetterRoutingKey, "string")
	addNonEmpty(arguments, "x-max-length", self.XMaxLength, "int32")
	addNonEmpty(arguments, "x-max-length-bytes", self.XMaxLengthBytes, "int32")
	addNonEmpty(arguments, "x-max-priority", self.XMaxPriority, "int32")
	addNonEmpty(arguments, "x-queue-mode", self.XQueueMode, "string")
	addNonEmpty(arguments, "x-queue-master-locator", self.XQueueMasterLocator, "string")

	return arguments
}

func addNonEmpty(arguments amqp.Table, key string, value interface{}, valueType string) {
	switch valueType {
	case "int32":
		if value.(int32) != 0 {
			arguments[key] = value
		}
	case "string":
		if value.(string) != "" {
			arguments[key] = value
		}
	case "bool":
		if value.(bool) {
			arguments[key] = value
		}
	}
}
