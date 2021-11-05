package main

import (
	"context"
	"fmt"
	"juegocalamar/pb"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

//type pozoMonto struct{
//	totalMonto int32
//}

var totalpozo int = 0

type LiderServer struct {
	pb.UnimplementedJugadorLiderServiceServer

	//Manejo de jugadores
	TotalPlayers      int32
	cantidadJugadores int32
	savedJugadores    [16]*pb.Jugador

	//jugadores que han muerto
	deadJugadores       [16]*pb.Jugador
	AlivePlayersCounter int32

	// info de las jugadas
	CurrentEtapa  int32
	CurrentRonda  int32
	TriggerReady  bool
	JugadasEtapas [3]jugadasRondas
}
type jugadasRondas struct {
	Etapa         int32
	ronda         int32
	jugadasRondas [3]jugadas
	total         int
}

type jugadas struct {
	jugadas       [16]jugada
	jugadaLider   jugada
	jugadasHechas int
}
type jugada struct {
	idPlayer  int
	playerObj *pb.Jugador
	Movement  int
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
		server.AlivePlayersCounter = server.AlivePlayersCounter + 1
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
		server.AlivePlayersCounter = server.AlivePlayersCounter + 1
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
	///*
	var Confirmacion string
	for {
		fmt.Printf("Iniciar Etapa1 [Y]/[N]: ")
		_, err := fmt.Scanf("%s", &Confirmacion)
		if err == nil {
			break
		}
	}
	if Confirmacion == "y" {
		if err := stream.Send(&pb.EsperandoJugadores{
			Message:      "Empezaremos a jugar",
			Confirmacion: true,
		}); err != nil {
			return err
		}
		return nil
	}
	//*/
	return nil

}

///*
func (server *LiderServer) LuzRojaLuzVerde(req *pb.JugadaCliente, stream pb.JugadorLiderService_LuzRojaLuzVerdeServer) error {
	var r1 int
	jugadaCl := jugada{
		idPlayer:  int(req.GetId()),
		playerObj: req.GetJugador(),
		Movement:  int(req.GetMessage()),
	}
	server.JugadasEtapas[server.CurrentEtapa].jugadasRondas[server.CurrentRonda].jugadas[jugadaCl.idPlayer] = jugadaCl
	server.JugadasEtapas[server.CurrentEtapa].jugadasRondas[server.CurrentRonda].jugadasHechas = server.JugadasEtapas[server.CurrentEtapa].jugadasRondas[server.CurrentRonda].jugadasHechas + 1
	log.Printf("Juegada recibida por jugador %d", &jugadaCl.idPlayer)
	if int(server.TotalPlayers) == server.JugadasEtapas[server.CurrentEtapa].jugadasRondas[server.CurrentRonda].jugadasHechas {
		log.Println("juegan los bots")
		for _, c := range server.savedJugadores {
			if c != nil {
				if c.Bot {
					if c.Vive {
						play := crearJugadaBot(c, 0, 10)
						server.JugadasEtapas[server.CurrentEtapa].jugadasRondas[server.CurrentRonda].jugadas[play.idPlayer] = play
						server.JugadasEtapas[server.CurrentEtapa].jugadasRondas[server.CurrentRonda].jugadasHechas = server.JugadasEtapas[server.CurrentEtapa].jugadasRondas[server.CurrentRonda].jugadasHechas + 1
						log.Printf("Juegada recibida por Bot %d", &play.idPlayer)
					}
				}
			}
		}
		log.Print(server.AlivePlayersCounter)
		log.Print(server.JugadasEtapas[server.CurrentEtapa].jugadasRondas[server.CurrentRonda].jugadasHechas)
	}
	if int(server.AlivePlayersCounter) == server.JugadasEtapas[server.CurrentEtapa].jugadasRondas[server.CurrentRonda].jugadasHechas {
		log.Println("Ahora juega el Lider")
		r1 = rand.Intn(4) + 6
		jugadaLid := jugada{
			idPlayer: -1,
			Movement: r1,
		}
		fmt.Print("Jugada del Lider: ")
		fmt.Printf("%d", r1)
		fmt.Println()
		server.JugadasEtapas[server.CurrentEtapa].Etapa = server.CurrentEtapa
		server.JugadasEtapas[server.CurrentEtapa].ronda = server.CurrentRonda
		server.JugadasEtapas[server.CurrentEtapa].jugadasRondas[server.CurrentRonda].jugadaLider = jugadaLid
		for _, c := range server.savedJugadores {
			if c != nil {
				if c.Vive {
					id := c.GetId()
					play := server.JugadasEtapas[server.CurrentEtapa].jugadasRondas[server.CurrentRonda].jugadas[id]
					if play.Movement > r1-1 {
						server.savedJugadores[c.Id].Vive = false
						server.deadJugadores[c.Id] = c
						fmt.Print("jugador Eliminado: ")
						fmt.Printf("%d", c.Id)
						fmt.Println()
					}

				}
			}
		}
		if server.CurrentRonda < 3 {
			stream.SendMsg(&pb.JugadaLider{
				Message:    int32(r1),
				ReadyEtapa: false,
			})
			server.CurrentRonda = server.CurrentRonda + 1
		} else {
			stream.SendMsg(&pb.JugadaLider{
				Message:    int32(r1),
				ReadyEtapa: true,
			})
			server.CurrentEtapa = server.CurrentEtapa + 1
			server.CurrentRonda = 1
		}
	}

	return nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func existeError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func crearArchivo() {
	var path = "pozo/pozo.txt"

	//Verifica que el archivo existe
	var _, err = os.Stat(path)
	//Crea el archivo si no existe
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if existeError(err) {
			return
		}
		defer file.Close()
	}
	//fmt.Println("File Created Successfully", path)
}

func EnviarAPozoJugadorEliminado(id int, n_ronda int, monto int) {
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

	var id_jugador string
	var id_ronda string
	var id_monto string
	id_jugador = strconv.Itoa(id)
	id_ronda = strconv.Itoa(n_ronda)
	id_monto = strconv.Itoa(monto) //el monto ya viene con el nuevo total

	body := "Jugador_" + id_jugador + " Ronda_" + id_ronda + " " + id_monto
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

func crearJugadaBot(jugador *pb.Jugador, floor int, top int) jugada {
	r1 := rand.Intn(top-floor) + floor
	jugadaBot := jugada{
		idPlayer:  int(jugador.GetId()),
		playerObj: jugador,
		Movement:  r1,
	}
	return jugadaBot
}

func main() {
	//totalpozo = totalpozo + 100000000
	//EnviarAPozoJugadorEliminado(3, 2, totalpozo) //Ejemplo para enviar pozo a jugador eliminado
	//EnviarAPozoJugadorEliminado(3, 2) //Ejemplo para enviar pozo a jugador eliminado

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

	pb.RegisterJugadorLiderServiceServer(s, &LiderServer{
		cantidadJugadores:   0,
		TotalPlayers:        int32(playersNo),
		AlivePlayersCounter: 0,
		CurrentEtapa:        1,
		CurrentRonda:        1,
	})
	reflection.Register(s)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	//Jugar
	//pb.Etapa1Server(s)

}
