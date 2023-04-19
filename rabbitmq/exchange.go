package rabbitmq

import log "github.com/sirupsen/logrus"

// DeclareExchange declares an exchange
func (self *Declarator) DeclareExchange(exchange Exchange) {
	err := self.conn.ExchangeDeclare(
		exchange.Name,
		exchange.Type,
		exchange.Durable,
		exchange.AutoDelete,
		exchange.Internal,
		exchange.NoWait,
		exchange.Arguments.GetArguments(),
	)
	if err != nil {
		log.Error("[RabbitMQ] [Exchange] Error when trying declare exchange " + err.Error())
	}

	log.Warn("[RabbitMQ] [Exchange] " + exchange.Name + " declared")
}

// DeclareExchanges declares all exchanges from a list of exchanges
func (self *Declarator) DeclareExchanges(exchange []Exchange) {
	for _, exchange := range exchange {
		self.DeclareExchange(exchange)
	}
}
