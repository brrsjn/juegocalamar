/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"juegocalamar/pb"

	"google.golang.org/grpc"
)

const (
	address    = "localhost:50051"
	defaultBot = true
)

/*type Jugada struct {
	Jugador *pb.Jugador
	mensaje string
	etapa   int32
}*/

func EsperarLider(client pb.JugadorLiderServiceClient, self *pb.Jugador) {

	stream, err := client.IniciarEtapa(context.TODO(), &pb.SolicitarInicioJuego{
		Id: self.Id,
	})
	if err != nil {
		log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
	}
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
		}
		fmt.Printf("%s", message.Message)

	}
	log.Printf("Listos Para jugar")
	EmpezaraJugarEtapa(client, self, 1)

}

func EmpezaraJugarEtapa(client pb.JugadorLiderServiceClient, self *pb.Jugador, etapa int32) {
	self = checkPlayerStatus(client, self)
	if self.Vive {
		if etapa == 1 {
			EnviarJugadaEtapa1(client, self)
		}
		if etapa == 2 {
			log.Print("Juego 2")
			//EnviarJugadaEtapa1(client, self)
		}
		if etapa == 3 {
			log.Print("Juego 3")
			//EnviarJugadaEtapa1(client, self)
		}
	} else {
		fmt.Println("El jugador ha sido eliminado")
	}
}

func checkPlayerStatus(client pb.JugadorLiderServiceClient, self *pb.Jugador) *pb.Jugador {
	self, _ = client.EstadoDelJugador(context.TODO(), self)
	return self
}

func EnviarJugadaEtapa1(client pb.JugadorLiderServiceClient, self *pb.Jugador) {

	//
	//reader := bufio.NewReader(os.Stdin)
	fmt.Println("Jugaremos,Muevete luz verde")
	var Movement int32
	for {
		fmt.Printf("pasos: ")
		_, err := fmt.Scanf("%d", &Movement)
		if err == nil {
			fmt.Println("JUGADA HECHA")
			self.SumaJugada1 = self.SumaJugada1 + Movement
			break
		}

	}
	stream, err := client.LuzRojaLuzVerde(context.TODO(), &pb.JugadaCliente{Id: self.GetId(), Message: Movement})
	if err != nil {
		log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
	}
	var NextEtapa int
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
		}
		log.Printf("No te muevas... (comparando %d con %d) ", message.Message, Movement)
		if message.ReadyEtapa {
			NextEtapa = 2
		} else {
			NextEtapa = 1
		}
	}
	EmpezaraJugarEtapa(client, self, int32(NextEtapa))
}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewJugadorLiderServiceClient(conn)

	// Contact the server and print out its response.
	bot := defaultBot
	if os.Args[len(os.Args)-1] == "--cli" {
		bot = false
	} else {
		bot = true
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SolicitarUnirce(ctx, &pb.InscripcionJugador{Bot: bot})
	if err != nil {
		log.Fatalf("could not greet: %v", err)

	}
	if r.Id == -1 {
		log.Printf("La sala esta llena")
	} else {
		log.Printf("Tu identificador es: %d", r.GetId())
		EsperarLider(c, r)
	}

}
