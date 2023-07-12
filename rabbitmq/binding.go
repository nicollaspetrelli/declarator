package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
)

// DeclareBindings declares all bindings from a list of bindings
func (d *Declarator) DeclareBindings(bindings []Binding) {
	for _, binding := range bindings {
		d.Bind(binding)
	}
}

// Bind binds a queue or exchange to another exchange
func (d *Declarator) Bind(binding Binding) {
	if binding.DestinationType == "queue" {
		d.bindQueue(binding)
	}

	if binding.DestinationType == "exchange" {
		d.bindExchange(binding)
	}
}

func (d *Declarator) bindQueue(binding Binding) {
	err := d.conn.QueueBind(
		binding.Destination,
		binding.RoutingKey,
		binding.Source,
		binding.NoWait,
		amqp.Table{},
	)
	if err != nil {
		log.Error("[RabbitMQ] [Binding] Error when trying bind queue " + err.Error())
		return
	}

	log.Warn("[RabbitMQ] [Binding] Queue " + binding.Destination + " binded with exchange " + binding.Source)
}

func (d *Declarator) bindExchange(binding Binding) {
	err := d.conn.ExchangeBind(
		binding.Destination,
		binding.RoutingKey,
		binding.Source,
		binding.NoWait,
		amqp.Table{},
	)
	if err != nil {
		log.Error("[RabbitMQ] [Binding] Error when trying bind exchange " + err.Error())
		return
	}

	log.Warn("[RabbitMQ] [Binding] Exchange " + binding.Destination + " binded with exchange " + binding.Source)
}
