package main

import (
	"fmt"

	consumer "github.com/gabrielpires-1/natural-events-alert-system/consumer/main"
	log "github.com/gabrielpires-1/natural-events-alert-system/log/main"
	producer "github.com/gabrielpires-1/natural-events-alert-system/producer/main"
)

func main() {
	fmt.Println("Welcome to the NEAS, the Natural-Events Alert System! Choose an option:")
	fmt.Println("1. Enter as producer")
	fmt.Println("2. Enter as consumer")
	fmt.Println("3. Enter as log")

	var option int
	fmt.Scanln(&option)

	switch option {
	case 1:
		producer.Run()
	case 2:
		consumer.Run()
	case 3:
		log.Run()
	default:
		fmt.Println("Invalid option")
	}
}
