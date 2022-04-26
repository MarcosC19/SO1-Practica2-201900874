package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	pb "github.com/MarcosC19/SO1-Practica2-201900874/grpc-client/protos"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type DataGame struct {
	GameId  int32 `json:"game_id"`
	Players int32 `json:"players"`
}

func PlayGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// GETTING DATA BODY FOR DATAGAME STRUCT
	var game DataGame

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &game)

	// CONNETING TO gRPC SERVER
	path := os.Getenv("IP_SERVER")
	if len(path) == 0 {
		path = "localhost:50051"
	}
	conn, err := grpc.Dial(path, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewPlayGameClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// CALLING THE METHOD PLAYING TO RUN A GAME
	reply, err := c.Playing(ctx, &pb.GameRequest{
		GameId:  game.GameId,
		Players: game.Players,
	})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	// RETURN THE RESPONSE
	json.NewEncoder(w).Encode(struct {
		Status int32 `json:"status"`
	}{Status: reply.GetStatus()})
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(struct {
		Response string `json:"response"`
	}{Response: "Server running"})
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Home)
	router.HandleFunc("/runGame", PlayGame).Methods("POST")
	log.Println("Listening at port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
