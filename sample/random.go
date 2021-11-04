package sample

import (
	"math/rand"
)

func randomNumberLiderEtapa1() int {
	r1 := rand.Intn(4) + 6
	return r1
}

func randomNumberJugadorEtapa1() int {
	return rand.Intn(10)
}
