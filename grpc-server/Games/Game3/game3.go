package Game3

import (
	"math/rand"
	"time"

	"github.com/MarcosC19/SO1-Practica2-201900874/grpc-server/Games/helpers"
)

// JUEGO NUMERO MAS GRANDE
func BigBrother(players int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	var playerList []int
	for i := 1; i <= players; i++ {
		playerList = append(playerList, i)
	}
	var list = helpers.PairPlayers(playerList, helpers.MAX)

	var maxium = helpers.Result{
		Player:    0,
		Selection: 0,
	}

	for _, s := range list {
		if s.P1.Selection > maxium.Selection {
			maxium = s.P1
		}

		if s.P2.Selection > maxium.Selection {
			maxium = s.P2
		}
	}

	return maxium.Player
}
