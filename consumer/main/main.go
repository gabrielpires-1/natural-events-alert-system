package consumer

import (
	"fmt"

	"github.com/gabrielpires-1/natural-events-alert-system/pubsub"
	amqp "github.com/rabbitmq/amqp091-go"
)

const exchange = "alert_topic"
const tempKey = "temperature"
const pressureKey = "pressure"
const rainKey = "rain"
const sismicKey = "sismic"

func Run() {
	fmt.Println("Starting system server...")
	string_connect := "amqp://guest:guest@localhost:5672/"
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

	fmt.Println("You are the consommer. Which metric you want to subscribe?")
	fmt.Println("1-temperature")
	fmt.Println("2-pressure")
	fmt.Println("3-rain")
	fmt.Println("4-sismic activity")

	var option int
	fmt.Scanln(&option)

	switch option {
	case 1:
		queue := ""
		pubsub.SubscribeJSON(connection, exchange, queue, "sensor."+tempKey, pubsub.SimpleQueueTransient, func(temp int) pubsub.AckType {
			fmt.Printf("Temperature: %d\n", temp)
			return pubsub.Ack
		})
		for {
		}
	case 2:
		queue := ""
		pubsub.SubscribeJSON(connection, exchange, queue, "sensor."+pressureKey, pubsub.SimpleQueueTransient, func(pressure int) pubsub.AckType {
			fmt.Printf("Pressure: %d\n", pressure)
			return pubsub.Ack
		})
		for {
		}
	case 3:
		queue := ""
		pubsub.SubscribeJSON(connection, exchange, queue, "sensor."+rainKey, pubsub.SimpleQueueTransient, func(rain int) pubsub.AckType {
			fmt.Printf("Rain: %d\n", rain)
			return pubsub.Ack
		})
		for {
		}
	case 4:
		queue := ""
		pubsub.SubscribeJSON(connection, exchange, queue, "sensor."+sismicKey, pubsub.SimpleQueueTransient, func(sismic int) pubsub.AckType {
			fmt.Printf("Sismic activity: %d\n", sismic)
			return pubsub.Ack
		})
		for {
		}
	}
}
