package Game1

import (
	"math/rand"
	"time"

	"github.com/MarcosC19/SO1-Practica2-201900874/grpc-server/Games/helpers"
)

// JUEGO PIEDRA PAPEL O TIJERA
func Rps(Players int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	var PlayerList []int
	for i := 1; i <= Players; i++ {
		PlayerList = append(PlayerList, i)
	}
	var list = helpers.PairPlayers(PlayerList, 3)
	var winner = helpers.ProcessPairsRPS(list)
	return winner
}
