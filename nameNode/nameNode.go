package main

import (
	"juegocalamar/pb"
	"log"
)

type NameNodeServer struct {
	pb.UnimplementedLiderNameNodeServiceServer
	jugadas []*pb.Jugada
}

func main() {
	log.Printf("NameNode iniciado, aqui se guardarán y se distribuiran los datos")
}
