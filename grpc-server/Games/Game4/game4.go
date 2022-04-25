package Game4

import (
	"math/rand"
	"time"

	"github.com/MarcosC19/SO1-Practica2-201900874/grpc-server/Games/helpers"
)

// JUEGO NUMERO MAS PEQUEÑO
func SmallBrother(players int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	var playerList []int
	for i := 1; i <= players; i++ {
		playerList = append(playerList, i)
	}
	var list = helpers.PairPlayers(playerList, helpers.MAX)

	var minium = helpers.Result{
		Player:    0,
		Selection: 101,
	}

	for _, s := range list {
		if s.P1.Selection < minium.Selection {
			minium = s.P1
		}

		if s.P2.Selection < minium.Selection {
			minium = s.P2
		}
	}

	return minium.Player
}
