package log

import (
	"fmt"
	producer "github.com/gabrielpires-1/natural-events-alert-system/producer/main"
	"github.com/gabrielpires-1/natural-events-alert-system/pubsub"
	amqp "github.com/rabbitmq/amqp091-go"
)

func Run(string_connect string) {
	fmt.Println("Starting system server...")

	connection, err := amqp.Dial(string_connect)
	if err != nil {
		fmt.Println("Failed to connect to RabbitMQ")
		panic(err)
	}
	defer connection.Close()

	fmt.Println("Connected to RabbitMQ")

	exchange := "alert_topic"
	queueName := "log_queue"
	routingKey := "sensor.#"

	publishCh, queue, err := pubsub.DeclareAndBind(connection, exchange, queueName, routingKey, pubsub.SimpleQueueDurable)
	if err != nil {
		fmt.Println("Failed to declare and bind queue")
		panic(err)
	}
	defer publishCh.Close()

	fmt.Printf("Queue %s is ready to receive messages\n", queue.Name)

	pubsub.SubscribeJSON(connection, exchange, queue.Name, routingKey, pubsub.SimpleQueueDurable, func(msg producer.Msg) pubsub.AckType {
		fmt.Printf(">Message: %v\n", msg)
		return pubsub.Ack
	})

	select {}
}
