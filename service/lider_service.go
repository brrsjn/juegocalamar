package main

import (
	"context"
	"juegocalamar/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type LiderServer struct {
	pb.UnimplementedJugadorLiderServiceServer
	savedJugadores    [16]*pb.Jugador
	cantidadJugadores int32
}

func viveJugador(lider int, jugador int) bool {
	if lider > jugador {
		return true
	} else {
		return false
	}
}

func NewLiderServer() *LiderServer {
	return &LiderServer{}
}

func (server *LiderServer) SolicitarUnirce(ctx context.Context, req *pb.InscripcionJugador) (*pb.Jugador, error) {
	log.Printf("Received bot: %v", req.GetBot())
	jugador := &pb.Jugador{
		Id:   server.cantidadJugadores,
		Bot:  req.GetBot(),
		Vive: true,
	}
	server.savedJugadores[server.cantidadJugadores] = jugador
	server.cantidadJugadores = server.cantidadJugadores + 1
	return jugador, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterJugadorLiderServiceServer(s, &LiderServer{cantidadJugadores: 0})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
