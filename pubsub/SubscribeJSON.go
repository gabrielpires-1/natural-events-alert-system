package pubsub

import (
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

func SubscribeJSON[T any](
	conn *amqp.Connection,
	exchange,
	queueName,
	key string,
	simpleQueueType int,
	handler func(T) AckType,
) error {
	ch, queue, err := DeclareAndBind(conn, exchange, queueName, key, simpleQueueType)
	if err != nil {
		return err
	}

	err = ch.Qos(10, 0, false)
	if err != nil {
		return err
	}

	deliveryCh, err := ch.Consume(queue.Name, "", false, false, false, false, nil)
	if err != nil {
		return err
	}

	go func() {
		for msg := range deliveryCh {
			var data T
			err := json.Unmarshal(msg.Body, &data)
			if err != nil {
				fmt.Println("Error unmarshalling message:", err)
				continue
			}
			ack := handler(data)
			if ack == Ack {
				msg.Ack(false)
			} else if ack == NackRequeue {
				msg.Nack(false, true)
			} else if ack == NackDiscard {
				msg.Nack(false, false)
			} else {
				fmt.Println("Invalid ack:", ack)
			}
		}
	}()
	return nil
}
