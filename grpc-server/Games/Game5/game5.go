package Game5

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func startRoulette(players []int) []int {
	var outPlayer = int(math.Round(rand.Float64() * float64(len(players)-1)))

	return append(players[:outPlayer], players[outPlayer+1:]...)
}

// JUEGO RULETA
func Roulette(players int) int {
	fmt.Println("Starting roulette for", players, "players")
	rand.Seed(time.Now().UTC().UnixNano())
	var playerList []int
	for i := 1; i <= players; i++ {
		playerList = append(playerList, i)
	}

	for len(playerList) > 1 {
		playerList = startRoulette(playerList)
	}

	var winner = playerList[0]

	return winner
}
