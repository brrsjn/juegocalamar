package main

import (
	"context"
	"fmt"
	"juegocalamar/pb"
	"log"
	"net"
	"time"

	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type LiderServer struct {
	pb.UnimplementedJugadorLiderServiceServer

	//Manejo de jugadores
	TotalPlayers      int32
	cantidadJugadores int32
	savedJugadores    [16]*pb.Jugador

	//jugadores que han muerto
	deadJugadores [16]*pb.Jugador

	// info de las jugadas
	EtapaNo       int32
	ronda         int32
	JugadasRondas [3]jugadas
}

type jugadas struct {
	Etapa   int
	ronda   int
	jugadas [16]jugada
	total   int
}
type jugada struct {
	idPlayer int
	Movement int
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
			Id:   server.cantidadJugadores,
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
	for server.cantidadJugadores < 15 {

		server.cantidadJugadores = server.cantidadJugadores + 1
		jugador := &pb.Jugador{
			Id:   server.cantidadJugadores,
			Bot:  true,
			Vive: true,
		}
		server.savedJugadores[server.cantidadJugadores] = jugador
		log.Printf("New bot player %d", jugador.Id)
	}
	if err := stream.Send(&pb.EsperandoJugadores{
		Message:      "Se han completado los jugadores",
		Confirmacion: true,
	}); err != nil {
		return err
	}
	return nil
}

///*
func (server *LiderServer) LuzRojaLuzVerde(req *pb.JugadaCliente, stream pb.JugadorLiderService_LuzRojaLuzVerdeServer) error {
	stream.SendMsg(&pb.JugadaLider{
		Message: 2,
	})

	return nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func EnviarAPozoJugadorEliminado(id int32, etapa int32) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := "Hello World!"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}

func main() {
	EnviarAPozoJugadorEliminado(3, 2) //Ejemplo para enviar pozo a jugador eliminado

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

	pb.RegisterJugadorLiderServiceServer(s, &LiderServer{cantidadJugadores: 0, TotalPlayers: int32(playersNo)})
	reflection.Register(s)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	//Jugar
	//pb.Etapa1Server(s)

}
