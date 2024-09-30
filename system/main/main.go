package main

import (
	"fmt"
	"os"
	consumer "github.com/gabrielpires-1/natural-events-alert-system/consumer/main"
	log "github.com/gabrielpires-1/natural-events-alert-system/log/main"
	producer "github.com/gabrielpires-1/natural-events-alert-system/producer/main"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar arquivo .env")
		os.Exit(1)
	}

	link := os.Getenv("AMQP_CONNECTION")
	if link == "" {
		fmt.Println("A variável AMQP_LINK não está definida no .env")
		os.Exit(1)
	}

	fmt.Println("Welcome to the NEAS, the Natural-Events Alert System! Choose an option:")
	fmt.Println("1. Enter as producer")
	fmt.Println("2. Enter as consumer")
	fmt.Println("3. Enter as log")

	var option int
	fmt.Scanln(&option)

	switch option {
	case 1:
		producer.Run(link)
	case 2:
		consumer.Run(link)
	case 3:
		log.Run(link)
	default:
		fmt.Println("Invalid option")
	}
}
