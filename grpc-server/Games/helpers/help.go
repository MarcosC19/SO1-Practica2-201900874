package helpers

import (
	"math"
	"math/rand"
)

const (
	PAPER    = 1
	ROCK     = 2
	SCISSORS = 3
	CARA     = 1
	ESCUDO   = 2
	MIN      = 1
	MAX      = 100
)

// STRUCT PARA CREAR LAS PAREJAS DE JUEGO
type PairStruct struct {
	P1 Result `json:"P1"`
	P2 Result `json:"P2"`
}

// STRUCT PARA CREAR LA SELECCION DE ELEMENTO
type Result struct {
	Player    int `json:"Player"`
	Selection int `json:"Selection"`
}

// SELECCION DE ELEMENTO
func RandomSel(Player int, max int) Result {
	ResultSel := Result{
		Player:    Player,
		Selection: int(math.Round(rand.Float64()*(float64(max)-1.0) + 1.0)),
	}

	return ResultSel
}

// SELECCION DE PAREJAS DE JUEGO
func PairPlayers(Players []int, max int) []PairStruct {
	var pairList []PairStruct
	for i := 0; i < len(Players); i = i + 2 {
		if i+1 < len(Players) {
			var pair = PairStruct{
				P1: RandomSel(Players[i], max),
				P2: RandomSel(Players[i+1], max),
			}
			pairList = append(pairList, pair)
		} else {
			var pair = PairStruct{
				P1: RandomSel(i+1, max),
				P2: RandomSel(i+1, max),
			}
			pairList = append(pairList, pair)
		}
	}
	return pairList
}

// SELECCION DEL GANADOR DE LA PAREJA JUEGO 1
func GetPairWinnerRPS(pair PairStruct) int {
	if pair.P1.Selection == PAPER && pair.P2.Selection == ROCK {
		return pair.P1.Player
	} else if pair.P1.Selection == PAPER && pair.P2.Selection == SCISSORS {
		return pair.P2.Player
	} else if pair.P1.Selection == PAPER && pair.P2.Selection == PAPER {
		if pair.P1.Player > pair.P2.Player {
			return pair.P1.Player
		} else {
			return pair.P2.Player
		}
	} else if pair.P1.Selection == ROCK && pair.P2.Selection == ROCK {
		if pair.P1.Player > pair.P2.Player {
			return pair.P1.Player
		} else {
			return pair.P2.Player
		}
	} else if pair.P1.Selection == ROCK && pair.P2.Selection == SCISSORS {
		return pair.P1.Player
	} else if pair.P1.Selection == ROCK && pair.P2.Selection == PAPER {
		return pair.P2.Player
	} else if pair.P1.Selection == SCISSORS && pair.P2.Selection == ROCK {
		return pair.P2.Player
	} else if pair.P1.Selection == SCISSORS && pair.P2.Selection == SCISSORS {
		if pair.P1.Player > pair.P2.Player {
			return pair.P1.Player
		} else {
			return pair.P2.Player
		}
	} else if pair.P1.Selection == SCISSORS && pair.P2.Selection == PAPER {
		return pair.P1.Player
	}
	return -1
}

// PROCESO PARA ELEGIR GANADOR FINAL JUEGO 1
func ProcessPairsRPS(pairArray []PairStruct) int {
	var winners []int
	for _, s := range pairArray {
		var winner = GetPairWinnerRPS(s)
		winners = append(winners, winner)
	}
	if len(winners) == 1 {
		return winners[0]
	}
	return ProcessPairsRPS(PairPlayers(winners, 3))
}

// SELECCION DE GANADOR DE PAREJA JUEGO 2
func GetPairWinnerFlip(pair PairStruct) int {
	var flipResult = int(math.Round(rand.Float64()*1.0 + 1.0))
	if pair.P1.Selection == flipResult {
		return pair.P1.Player
	}
	return pair.P2.Player
}

// PROCESO PARA ELEGIR GANADOR FINAL JUEGO 2
func ProcessPairsFlip(pairArray []PairStruct) int {
	var winners []int
	for _, s := range pairArray {
		var winner = GetPairWinnerFlip(s)
		winners = append(winners, winner)
	}
	if len(winners) == 1 {
		return winners[0]
	}
	return ProcessPairsFlip(PairPlayers(winners, 2))
}
