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

	publishCh, err := connection.Channel()
	if err != nil {
		fmt.Println("Failed to open a channel")
		panic(err)
	}
	defer publishCh.Close()

	fmt.Println("You are the log. Be ready to receive messages.")
	queue := ""
	exchange := "alert_topic"
	pubsub.SubscribeJSON(connection, exchange, queue, "sensor.#", pubsub.SimpleQueueTransient, func(msg producer.Msg) pubsub.AckType {
		fmt.Printf(">Message: %v\n", msg)
		return pubsub.Ack
	})
	for {
	}
}
