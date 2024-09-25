package consumer

import (
	"fmt"

	producer "github.com/gabrielpires-1/natural-events-alert-system/producer/main"
	"github.com/gabrielpires-1/natural-events-alert-system/pubsub"
	amqp "github.com/rabbitmq/amqp091-go"
)

const exchange = "alert_topic"
const tempKey = "temperature"
const pressureKey = "pressure"
const rainKey = "rain"
const sismicKey = "sismic"
const allKey = "*"

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

	fmt.Println("You are the consumer. What location you want to subscribe? (\"*\" for  all)")
	var location string
	fmt.Scanln(&location)

	fmt.Println("Which metric you want to subscribe?")
	fmt.Println("1-temperature")
	fmt.Println("2-pressure")
	fmt.Println("3-rain")
	fmt.Println("4-sismic activity")
	fmt.Println("5-all")

	var option int
	fmt.Scanln(&option)

	switch option {
	case 1:
		queue := ""
		pubsub.SubscribeJSON(connection, exchange, queue, "sensor."+location+"."+tempKey, pubsub.SimpleQueueTransient, func(msg producer.Msg) pubsub.AckType {
			fmt.Printf("--------------------------------------------------\n")
			fmt.Printf("Time: %s\nLocation: %s\nTopic: %s\nValue: %d\n", msg.Time, msg.Location, msg.Topic, msg.Value)
			if msg.Value > 40 {
				fmt.Printf("ALERT: Temperature above 40 degrees\n")
			} else if msg.Value < 0 {
				fmt.Printf("ALERT: Temperature below 10 degrees\n")
			}
			fmt.Printf("--------------------------------------------------\n")
			return pubsub.Ack
		})
		for {
		}
	case 2:
		queue := ""
		pubsub.SubscribeJSON(connection, exchange, queue, "sensor."+location+"."+pressureKey, pubsub.SimpleQueueTransient, func(msg producer.Msg) pubsub.AckType {
			fmt.Printf("--------------------------------------------------\n")
			fmt.Printf("Time: %s\nLocation: %s\nTopic: %s\nValue: %d\n", msg.Time, msg.Location, msg.Topic, msg.Value)
			if msg.Value > 1013 {
				fmt.Printf("ALERT: Pressure above 1013 hPa\n")
			} else if msg.Value < 1000 {
				fmt.Printf("ALERT: Pressure below 1000 hPa\n")
			}
			fmt.Printf("--------------------------------------------------\n")
			return pubsub.Ack
		})
		for {
		}
	case 3:
		queue := ""
		pubsub.SubscribeJSON(connection, exchange, queue, "sensor."+location+"."+rainKey, pubsub.SimpleQueueTransient, func(msg producer.Msg) pubsub.AckType {
			fmt.Printf("--------------------------------------------------\n")
			fmt.Printf("Time: %s\nLocation: %s\nTopic: %s\nValue: %d\n", msg.Time, msg.Location, msg.Topic, msg.Value)
			if msg.Value > 50 {
				fmt.Printf("ALERT: Rain above 50 mm\n")
			}
			fmt.Printf("--------------------------------------------------\n")
			return pubsub.Ack
		})
		for {
		}
	case 4:
		queue := ""
		pubsub.SubscribeJSON(connection, exchange, queue, "sensor."+location+"."+sismicKey, pubsub.SimpleQueueTransient, func(msg producer.Msg) pubsub.AckType {
			fmt.Printf("--------------------------------------------------\n")
			fmt.Printf("Time: %s\nLocation: %s\nTopic: %s\nValue: %d\n", msg.Time, msg.Location, msg.Topic, msg.Value)
			if msg.Value > 5 {
				fmt.Printf("ALERT: Sismic activity above 5\n")
			}
			fmt.Printf("--------------------------------------------------\n")
			return pubsub.Ack
		})
		for {
		}
	case 5:
		queue := ""
		pubsub.SubscribeJSON(connection, exchange, queue, "sensor."+location+"."+allKey, pubsub.SimpleQueueTransient, func(msg producer.Msg) pubsub.AckType {
			fmt.Printf("--------------------------------------------------\n")
			fmt.Printf("Time: %s\nLocation: %s\nTopic: %s\nValue: %d\n", msg.Time, msg.Location, msg.Topic, msg.Value)
			if msg.Topic == tempKey {
				if msg.Value > 40 {
					fmt.Printf("ALERT: Temperature above 40 degrees\n")
				} else if msg.Value < 0 {
					fmt.Printf("ALERT: Temperature below 10 degrees\n")
				}
			} else if msg.Topic == pressureKey {
				if msg.Value > 1013 {
					fmt.Printf("ALERT: Pressure above 1013 hPa\n")
				} else if msg.Value < 1000 {
					fmt.Printf("ALERT: Pressure below 1000 hPa\n")
				}
			} else if msg.Topic == rainKey {
				if msg.Value > 50 {
					fmt.Printf("ALERT: Rain above 50 mm\n")
				}
			} else if msg.Topic == sismicKey {
				if msg.Value > 5 {
					fmt.Printf("ALERT: Sismic activity above 5\n")
				}
			}
			fmt.Printf("--------------------------------------------------\n")
			return pubsub.Ack
		})
		for {
		}
	}
}
