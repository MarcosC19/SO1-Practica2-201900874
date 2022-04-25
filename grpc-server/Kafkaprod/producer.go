package kafkaprod

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

// JSON PARA ENVIAR A KAFKA
type Logs struct {
	Game_id   int32  `json:"game_id"`
	Players   int32  `json:"players_num"`
	Game_name string `json:"game_name"`
	Winner    int32  `json:"winner"`
	Queue     string `json:"queue"`
}

// PRODUCIENDO MENSAJES A LA COLA DE KAFKA
func SendKafka(id_game int32, num_players int32, name_game string, winner_game int32) {
	// CREANDO JSON
	myLog := Logs{
		Game_id:   id_game,
		Players:   num_players,
		Game_name: name_game,
		Winner:    winner_game,
		Queue:     "Apache Kafka",
	}

	// OBTENIENDO STRUCT EN FORMATO JSON
	jsonString, err := json.Marshal(myLog)

	if err != nil {
		fmt.Println("Ocurrio un error json: ", err)
	}

	// CONVIRTIENDO JSON EN STRING
	logString := string(jsonString)

	fmt.Println(logString)

	// CREATING PRODUCER KAFKA
	topic := "so1-proyecto-fase2"
	partition := 0

	host := os.Getenv("HOSTIP_KAFKA")

	if len(host) == 0 {
		host = "localhost:9092"
	}

	hostname := host

	fmt.Println(hostname)

	// CREANDO CONEXION
	conn, err := kafka.DialLeader(context.Background(), "tcp", hostname, topic, partition)
	if err != nil {
		log.Fatal("Ocurrio un error kafka: ", err.Error())
	}

	// RECORRIENDO TODO EL JSON EN STRING
	for _, word := range []string{string(logString)} {
		_, err = conn.WriteMessages(
			kafka.Message{Value: []byte(word)},
		)
		if err != nil {
			log.Fatal("Error al mandar el mensaje: ", err)
		}
	}

	if err := conn.Close(); err != nil {
		log.Fatal("Error al cerrar la conexion: ", err)
	}
}
