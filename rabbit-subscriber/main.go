package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo/readpref"

	_ "github.com/go-sql-driver/mysql"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	RABBIT_SERVER_ENV    = "RABIT_HOST"
	RABBIT_PORT_ENV      = "RABBIT_PORT"
	RABBIT_QUEUE_ENV     = "RABBIT_QUEUE"
	RABBIT_USER_ENV      = "RABBIT_USER"
	RABBIT_PASS_ENV      = "RABBIT_PASSWORD"
	MONGO_HOST_ENV       = "MONGO_HOST"
	MONGO_PORT_ENV       = "MONGO_PORT"
	MONGO_USER_ENV       = "MONGO_USER"
	MONGO_PASS_ENV       = "MONGO_PASS"
	MONGO_DB_ENV         = "MONGO_DB"
	MONGO_COLLECTION_ENV = "MONGO_COLLECTION"
)

var (
	RabbitHost      = getEnv(RABBIT_SERVER_ENV, "localhost")
	RabbitPort      = getEnv(RABBIT_PORT_ENV, "5672")
	RabbitQueue     = getEnv(RABBIT_QUEUE_ENV, "GameQueue")
	RabbitUser      = getEnv(RABBIT_USER_ENV, "rabbit")
	RabbitPass      = getEnv(RABBIT_PASS_ENV, "sopes1")
	MongoHost       = getEnv(MONGO_HOST_ENV, "192.168.1.10")
	MongoPort       = getEnv(MONGO_PORT_ENV, "27017")
	MongoUser       = getEnv(MONGO_USER_ENV, "admin201900874")
	MongoPass       = getEnv(MONGO_PASS_ENV, "so1-practica2")
	MongoDB         = getEnv(MONGO_DB_ENV, "so1-practica2")
	MongoCollection = getEnv(MONGO_COLLECTION_ENV, "logs")
)

type Log struct {
	Game_id   int32  `json:"game_id"`
	Players   int32  `json:"players"`
	Game_name string `json:"game_name"`
	Winner    int32  `json:"winner"`
	Queue     string `json:"queue"`
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func main() {
	// Wait for Rabbit to Start Delay
	time.Sleep(15 * time.Second)

	// Start the RabbitMQ connection using credentials
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s", RabbitUser, RabbitPass, RabbitHost, RabbitPort))

	if err != nil {
		fmt.Println("Error connecting to Rabbit", err)
		return
	}
	defer conn.Close()

	// Create the Channel
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("Error creating the channel", err)
		return
	}
	defer ch.Close()

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
		fmt.Printf("Error subscribing to %s queue, error: %s \n", RabbitQueue, err)
		return
	}

	fmt.Println("Connection succeed to RabbitMQ")
	fmt.Println("Waiting for messages...")

	go listenMessages(messages)
	select {}
}

func listenMessages(messages <-chan amqp.Delivery) {
	for message := range messages {
		var log Log
		fmt.Println("Raw Message: ", string(message.Body))
		err := json.Unmarshal(message.Body, &log)
		if err != nil {
			fmt.Println("Error marshalling", err)
			return
		}

		fmt.Println("##########################")
		fmt.Println("# New Message: ")
		fmt.Printf("# GameID: %d\n", log.Game_id)
		fmt.Printf("# Players: %d\n", log.Players)
		fmt.Printf("# GameName: %s\n", log.Game_name)
		fmt.Printf("# Winner: %d\n", log.Winner)
		fmt.Printf("# Queue: %s\n", log.Queue)
		fmt.Println("##########################")

		go saveToMongo(log)
	}
}
func saveToMongo(data Log) {
	client, err := mongo.Connect(
		context.Background(),
		options.Client().ApplyURI("mongodb://"+MongoUser+":"+MongoPass+"@"+MongoHost+":"+MongoPort),
	)

	if err != nil {
		fmt.Println("Error connecting to MongoDB: ", err)
		return
	}
	defer client.Disconnect(context.Background())

	err = client.Ping(context.Background(), readpref.Primary())

	if err != nil {
		fmt.Println("Error pinging connection with MongoDB: ", err)
		return
	}

	Collection := client.Database(MongoDB).Collection(MongoCollection)

	insertResult, err := Collection.InsertOne(context.Background(), data)

	if err != nil {
		fmt.Println("Error inserting data: ", err)
	}

	fmt.Println("Nuevo log insertado en Mongo: ", insertResult)
}
