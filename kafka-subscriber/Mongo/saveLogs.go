package Mongo

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	port = 27017
	//user     = "admingrupo19"
	//password = "so1-fase2"
)

// JSON PARA RECIBIR DE KAFKA
type Logs struct {
	Game_id   int32  `json:"game_id"`
	Players   int32  `json:"players_num"`
	Game_name string `json:"game_name"`
	Winner    int32  `json:"winner"`
	Queue     string `json:"queue"`
}

// GUARDANDO LOGS EN MONGODB
func SaveLogs(logsData Logs) {
	host := os.Getenv("HOSTIP_MONGO")

	if len(host) == 0 {
		host = "localhost"
	}

	// OPENING CONNECTION TO MONGODB
	clientOpts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", host, port))
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
	collection := client.Database("so-proyecto-f2").Collection("logs")

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
