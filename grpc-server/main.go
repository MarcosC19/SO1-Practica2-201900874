package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/MarcosC19/SO1-Practica2-201900874/grpc-server/Games/Game1"
	"github.com/MarcosC19/SO1-Practica2-201900874/grpc-server/Games/Game2"
	"github.com/MarcosC19/SO1-Practica2-201900874/grpc-server/Games/Game3"
	"github.com/MarcosC19/SO1-Practica2-201900874/grpc-server/Games/Game4"
	"github.com/MarcosC19/SO1-Practica2-201900874/grpc-server/Games/Game5"
	kafkaprod "github.com/MarcosC19/SO1-Practica2-201900874/grpc-server/Kafkaprod"
	pb "github.com/MarcosC19/SO1-Practica2-201900874/grpc-server/protos"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement server.PlayGameServer.
type server struct {
	pb.UnimplementedPlayGameServer
}

// calling the function rcp Playing
func (s *server) Playing(ctx context.Context, in *pb.GameRequest) (*pb.GameReply, error) {
	log.Printf("Received: { %v, %v }", in.GetGameId(), in.GetPlayers())
	var game_name string
	var winner int

	switch in.GetGameId() {
	case 1:
		game_name = "Piedra, papel o tijeras"
		winner = Game1.Rps(int(in.GetPlayers()))
	case 2:
		game_name = "Cara o Cruz"
		winner = Game2.Flip(int(in.GetPlayers()))
	case 3:
		game_name = "Numero mas grande"
		winner = Game3.BigBrother(int(in.GetPlayers()))
	case 4:
		game_name = "Numero mas pequeño"
		winner = Game4.SmallBrother(int(in.GetPlayers()))
	case 5:
		game_name = "Ruleta"
		winner = Game5.Roulette(int(in.GetPlayers()))
	}

	kafkaprod.SendKafka(in.GameId, in.Players, game_name, int32(winner))

	return &pb.GameReply{GameId: in.GameId, Players: in.Players, GameName: game_name, Winner: int32(winner)}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPlayGameServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
