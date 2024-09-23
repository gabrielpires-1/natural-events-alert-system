package productor

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

	fmt.Println("You are the productor. Which metric you can share?")
	fmt.Println("1-temperature")
	fmt.Println("2-pressure")
	fmt.Println("3-rain")
	fmt.Println("4-sismic activity")

	var option int
	fmt.Scanln(&option)

	switch option {
	case 1:
		for {
			var temperature int
			fmt.Println("Input temperature:")
			fmt.Scanln(&temperature)
			pubsub.PublishJSON(publishCh, exchange, "sensor."+tempKey, temperature)
		}
	case 2:
		for {
			var pressure int
			fmt.Println("Input pressure:")
			fmt.Scanln(&pressure)
			pubsub.PublishJSON(publishCh, exchange, "sensor."+pressureKey, pressure)
		}
	case 3:
		for {
			var rain int
			fmt.Println("Input rain level:")
			fmt.Scanln(&rain)
			pubsub.PublishJSON(publishCh, exchange, "sensor."+rainKey, rain)
		}
	case 4:
		var sismic int
		fmt.Println("Input sismic level:")
		fmt.Scanln(&sismic)
		pubsub.PublishJSON(publishCh, exchange, "sensor."+sismicKey, sismic)
	}
}
