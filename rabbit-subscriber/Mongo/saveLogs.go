package Mongo

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	port     = 27017
	user     = "admin201900874"
	password = "so1-practica2"
)

// JSON PARA RECIBIR DE RABBIT
type Logs struct {
	Request_number int32  `json:"request_number"`
	Game_id        int32  `json:"game_id"`
	Players        int32  `json:"players"`
	Game_name      string `json:"game_name"`
	Winner         int32  `json:"winner"`
	Queue          string `json:"queue"`
}

func CountLogs() int32 {
	host := os.Getenv("HOSTIP_MONGO")

	if len(host) == 0 {
		host = "localhost"
	}

	// OPENING CONNECTION TO MONGODB
	clientOpts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%d", user, password, host, port))
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}

	// VERIFYING THE CONNECTION
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexion exitosa")

	// CONNECTION TO DATABASE AND COLLECTION
	collection := client.Database("so1-practica2").Collection("logs")

	// GET OPERATIONS
	data, err := collection.CountDocuments(context.TODO(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	return int32(data)
}

// GUARDANDO LOGS EN MONGODB
func SaveLogs(logsData Logs) {
	logsData.Request_number = CountLogs() + 1

	host := os.Getenv("HOSTIP_MONGO")

	if len(host) == 0 {
		host = "localhost"
	}

	// OPENING CONNECTION TO MONGODB
	clientOpts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%d", user, password, host, port))
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}

	// VERIFYING THE CONNECTION
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexion exitosa")

	// CONNECTION TO DATABASE AND COLLECTION
	collection := client.Database("so1-practica2").Collection("logs")

	// INSERT THE NEW OPERATION
	insertResult, err := collection.InsertOne(context.TODO(), logsData)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Nuevo log insertado en Mongo con exito ", insertResult)

	//CLOSING CONNECTION TO MONGODB
	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexion cerrada")
}
