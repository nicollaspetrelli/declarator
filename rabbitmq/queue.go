package rabbitmq

import log "github.com/sirupsen/logrus"

func (self *Declarator) DeclareQueue(queue Queue) {
	declaredQueue, err := self.conn.QueueDeclare(
		queue.Name,
		queue.Durable,
		queue.AutoDelete,
		queue.Exclusive,
		queue.NoWait,
		queue.Arguments.GetArguments(),
	)
	if err != nil {
		log.Error("[RabbitMQ] [Queue] Error when trying declare queue " + err.Error())
		return
	}

	log.Warn("[RabbitMQ] [Queue] " + declaredQueue.Name + " declared")
}

func (self *Declarator) DeclareQueues(queues []Queue) {
	for _, queue := range queues {
		self.DeclareQueue(queue)
	}
}
