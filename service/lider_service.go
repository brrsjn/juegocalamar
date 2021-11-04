package service

import (
	"context"
	"juegocalamar/pb"
)

type LiderServer struct {
}

func NewLiderServer() *LiderServer {
	return &LiderServer{}
}

func (server *LiderServer) CreateLider(ctx context.Context, req *pb.Jugador) (*pb.Jugador, error) {
	solicitud := req.JugadorLiderService
}
