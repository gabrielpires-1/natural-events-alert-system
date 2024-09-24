package producer

import (
	"fmt"
	"strings"
	"time"

	"github.com/gabrielpires-1/natural-events-alert-system/pubsub"
	amqp "github.com/rabbitmq/amqp091-go"
)

const exchange = "alert_topic"
const tempKey = "temperature"
const pressureKey = "pressure"
const rainKey = "rain"
const sismicKey = "sismic"

// msg struct: [TIME] - [LOCALIZAÇÃO] - [TOPIC] - [VALUE]
type Msg struct {
	Time     string
	Location string
	Topic    string
	Value    int
}

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

	fmt.Println("Input location: ")
	var location string
	fmt.Scanln(&location)
	location = strings.ToLower(location)

	switch option {
	case 1:
		for {
			var temperature int
			fmt.Println("Input temperature:")
			fmt.Scanln(&temperature)
			msg := Msg{
				Time:     time.Now().Format("2006-01-02 15:04:05"),
				Location: location,
				Topic:    tempKey,
				Value:    temperature,
			}
			pubsub.PublishJSON(publishCh, exchange, "sensor."+location+"."+tempKey, msg)
		}
	case 2:
		for {
			var pressure int
			fmt.Println("Input pressure:")
			fmt.Scanln(&pressure)
			msg := Msg{
				Time:     time.Now().Format("2006-01-02 15:04:05"),
				Location: location,
				Topic:    pressureKey,
				Value:    pressure,
			}
			pubsub.PublishJSON(publishCh, exchange, "sensor."+location+"."+pressureKey, msg)
		}
	case 3:
		for {
			var rain int
			fmt.Println("Input rain level:")
			fmt.Scanln(&rain)
			msg := Msg{
				Time:     time.Now().Format("2006-01-02 15:04:05"),
				Location: location,
				Topic:    rainKey,
				Value:    rain,
			}
			pubsub.PublishJSON(publishCh, exchange, "sensor."+location+"."+rainKey, msg)
		}
	case 4:
		for {
			var sismic int
			fmt.Println("Input sismic level:")
			fmt.Scanln(&sismic)
			msg := Msg{
				Time:     time.Now().Format("2006-01-02 15:04:05"),
				Location: location,
				Topic:    sismicKey,
				Value:    sismic,
			}
			pubsub.PublishJSON(publishCh, exchange, "sensor."+location+"."+sismicKey, msg)
		}
	}
}
