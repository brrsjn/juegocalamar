package main

import (
	"context"
	"fmt"
	"juegocalamar/pb"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50063"
)

type DataNode3Server struct {
	pb.UnimplementedDataNameNodeServiceServer
}

func (server *DataNode3Server) RegistraJugada(ctx context.Context, req *pb.JugadaToDataNode) (*pb.Response, error) {
	log.Printf("Jugaba recibida para jugador: %d", req.GetId())
	path := fmt.Sprintf("nameNode/dataNode3/jugadas/jugador_%d_ronda_%d.txt", req.GetId(), req.GetEtapa())
	crearArchivo(path)
	escribeArchivo(path, int(req.GetMovimiento()))
	res := &pb.Response{
		State: true,
	}
	return res, nil
}

func crearArchivo(path string) {
	//Verifica que el archivo existe
	var _, err = os.Stat(path)
	//Crea el archivo si no existe
	if os.IsNotExist(err) {
		log.Printf("Hola")
		var file, err = os.Create(path)
		if existeError(err) {
			return
		}
		defer file.Close()
	}
	fmt.Println("File Created Successfully", path)
}

func escribeArchivo(path string, movimiento int) {
	// Abre archivo usando permisos READ & WRITE
	var file, err = os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0660)
	if existeError(err) {
		return
	}
	defer file.Close()
	// Escribe algo de texto linea por linea
	text := fmt.Sprintf("%d\n", movimiento)
	_, err = file.WriteString(text)
	if existeError(err) {
		return
	}
	// Salva los cambios
	err = file.Sync()
	if existeError(err) {
		return
	}
	fmt.Println("Archivo actualizado existosamente.")
}

func existeError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDataNameNodeServiceServer(s, &DataNode3Server{})
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
