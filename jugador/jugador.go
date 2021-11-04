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

type Jugada struct {
	Jugador *pb.Jugador
	mensaje string
	etapa   int32
}

func EsperarLider(client pb.JugadorLiderServiceClient, id int32) {
	stream, err := client.IniciarEtapa(context.TODO(), &pb.SolicitarInicioJuego{
		Id: id,
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
		log.Printf("%s", message.Message)
	}
	log.Printf("Listos Para jugar")
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
	if len(os.Args) > 1 {
		if os.Args[1] == "--cli" {
			bot = true
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SolicitarUnirce(ctx, &pb.InscripcionJugador{Bot: bot})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Tu identificador es: %d", r.GetId())
	EsperarLider(c, r.GetId())
}
