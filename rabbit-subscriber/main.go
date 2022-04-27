package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/MarcosC19/SO1-Practica2-201900874/rabbit-subscriber/Mongo"
	"github.com/streadway/amqp"
)

const (
	RabbitPort  = "5672"
	RabbitQueue = "GamesQueue"
	RabbitUser  = "rabbit"
	RabbitPass  = "sopes1"
)

func main() {
	time.Sleep(60 * time.Second)

	// GETTING ENVIROMENT "RABBIT_HOST"
	RabbitHost := os.Getenv("RABBIT_HOST")
	if len(RabbitHost) == 0 {
		RabbitHost = "localhost"
	}

	// CONNECTING TO RABBIT WITH CREDENTIALS
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s", RabbitUser, RabbitPass, RabbitHost))

	if err != nil {
		fmt.Println("Error al conectar a Rabbit", err)
		return
	}
	defer conn.Close()

	// CREATING THE CHANNEL TO CONNECT
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("Error al crear el canal", err)
		return
	}
	defer ch.Close()

	// GETTING MESSAGES FROM THE QUEUE
	messages, err := ch.Consume(
		RabbitQueue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Printf("Error al suscribirse a la cola %s : %s \n", RabbitQueue, err)
		return
	}

	fmt.Println("Conexion a Rabbit exitosa")

	go ReadRabbit(messages)
	select {}
}

func ReadRabbit(messages <-chan amqp.Delivery) {
	for message := range messages {
		fmt.Println("Data obtenida: ", string(message.Body))

		// GETTING DATA TO STRUCT LOGS
		var log Mongo.Logs
		err := json.Unmarshal(message.Body, &log)
		if err != nil {
			fmt.Println("Error marshalling", err)
			return
		}

		// SAVE THE LOGS INTO MONGO
		go Mongo.SaveLogs(log)
	}
}
