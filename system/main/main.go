package main

import (
	"fmt"

	consumer "github.com/gabrielpires-1/natural-events-alert-system/consumer/main"
	log "github.com/gabrielpires-1/natural-events-alert-system/log/main"
	productor "github.com/gabrielpires-1/natural-events-alert-system/productor/main"
)

func main() {
	fmt.Println("Welcome to the NEAS, the Natural-Events Alert System! Choose an option:")
	fmt.Println("1. Enter as productor")
	fmt.Println("2. Enter as consumer")
	fmt.Println("3. Enter as log")

	var option int
	fmt.Scanln(&option)

	switch option {
	case 1:
		productor.Run()
	case 2:
		consumer.Run()
	case 3:
		log.Run()
	default:
		fmt.Println("Invalid option")
	}
}
