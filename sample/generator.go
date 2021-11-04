package sample

import (
	"juegocalamar/pb"
)

func newJugador(isBot bool) *pb.Jugador {
	jugador := &pb.Jugador{
		Bot:  isBot,
		Vive: true,
	}
	return jugador
}
