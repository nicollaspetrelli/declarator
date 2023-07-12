package rabbitmq

import log "github.com/sirupsen/logrus"

// DeclareQueue declares a queue
func (d *Declarator) DeclareQueue(queue Queue) {
	declaredQueue, err := d.conn.QueueDeclare(
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

// DeclareQueues declares all queues from a list of queues
func (d *Declarator) DeclareQueues(queues []Queue) {
	for _, queue := range queues {
		d.DeclareQueue(queue)
	}
}
