package pubsub

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

func DeclareAndBind(
	conn *amqp.Connection,
	exchange,
	queueName,
	key string,
	simpleQueueType int, // an enum to represent "durable" or "transient"
) (*amqp.Channel, amqp.Queue, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, amqp.Queue{}, err
	}

	var durable, autoDelete, exclusive bool
	if simpleQueueType == SimpleQueueDurable {
		durable = true
		autoDelete = false
		exclusive = false
	} else if simpleQueueType == SimpleQueueTransient {
		durable = false
		autoDelete = true
		exclusive = true
	}

	var queue amqp.Queue
	queue, err = ch.QueueDeclare(
		queueName,
		durable,
		autoDelete,
		exclusive,
		false,
		nil)

	if err != nil {
		return nil, amqp.Queue{}, err
	}

	err = ch.QueueBind(
		queue.Name,
		key,
		exchange,
		false,
		nil)
	if err != nil {
		return nil, amqp.Queue{}, err
	}
	fmt.Printf("Queue %s bound to exchange %s with key %s\n", queue.Name, exchange, key)

	return ch, queue, nil
}
