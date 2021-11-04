package main

import (
	"context"
	"fmt"
	"juegocalamar/pb"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type LiderServer struct {
	pb.UnimplementedJugadorLiderServiceServer
	savedJugadores    [16]*pb.Jugador
	deadJugadores     [16]*pb.Jugador
	cantidadJugadores int32
	TotalPlayers      int32
	stageOneReady     bool
	stageTwoReady     bool
	stageThree        bool
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
	if server.cantidadJugadores < server.TotalPlayers {
		log.Printf("Received bot: %v", req.GetBot())
		jugador := &pb.Jugador{
			Id:   server.cantidadJugadores + 1,
			Bot:  req.GetBot(),
			Vive: true,
		}

		server.savedJugadores[server.cantidadJugadores] = jugador
		server.cantidadJugadores = server.cantidadJugadores + 1
		return jugador, nil
	} else {
		jugador := &pb.Jugador{
			Id:   -1,
			Bot:  req.GetBot(),
			Vive: false,
		}
		return jugador, nil
	}

}

func (server *LiderServer) IniciarEtapa(req *pb.SolicitarInicioJuego, stream pb.JugadorLiderService_IniciarEtapaServer) error {
	for server.cantidadJugadores < server.TotalPlayers {
		time.Sleep(1 * time.Second)
	}
	if err := stream.Send(&pb.EsperandoJugadores{
		Message: "Se han completado los jugadores",
	}); err != nil {
		return err
	}
	log.Printf("Inicio de etapa 1")
	//logica de inicio de etapa
	if err := stream.Send(&pb.EsperandoJugadores{
		Message: "Se dará inicio a la primera etapa \n",
	}); err != nil {
		return err
	}
	return nil
}

func main() {

	//Bienvenida e inicio del juego
	fmt.Println("___Bienvenido al juego del Calamardo___")
	fmt.Println("//deathmatch de 16 concursantes para ganar un premio suculento//")
	fmt.Print("Ingrese numero de jugadores (el resto seran bots): ")
	//reader := bufio.NewReader(os.Stdin)

	var playersNo int
	_, err := fmt.Scanf("%d", &playersNo)
	//if err != nil

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterJugadorLiderServiceServer(s, &LiderServer{cantidadJugadores: 0, TotalPlayers: int32(playersNo), stageOneReady: false})
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
