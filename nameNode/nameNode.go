package main

import (
	"context"
	"fmt"
	"juegocalamar/pb"
	"log"
	"math/rand"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
)

const (
	address1 = "localhost:50061"
	address2 = "localhost:50062"
	address3 = "localhost:50063"
	port     = ":50052"
)

type NameNodeServer struct {
	pb.UnimplementedLiderNameNodeServiceServer
	cliente1 pb.DataNameNodeServiceClient
	cliente2 pb.DataNameNodeServiceClient
	cliente3 pb.DataNameNodeServiceClient
}

func EnviarJugadaADataNode(cliente pb.DataNameNodeServiceClient, jugada *pb.JugadaToDataNode, pos int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	path := "nameNode/distribucion.txt"
	crearArchivo(path)
	if pos == 0 {
		texto := fmt.Sprintf("Jugador_%d_etapa_%d_%s", jugada.GetId(), jugada.GetEtapa(), address1)
		escribeArchivo(path, texto)
	}
	if pos == 1 {
		texto := fmt.Sprintf("Jugador_%d_etapa_%d_%s", jugada.GetId(), jugada.GetEtapa(), address2)
		escribeArchivo(path, texto)
	}
	if pos == 2 {
		texto := fmt.Sprintf("Jugador_%d_etapa_%d_%s", jugada.GetId(), jugada.GetEtapa(), address3)
		escribeArchivo(path, texto)
	}

	r, err := cliente.RegistraJugada(ctx, jugada)
	if err != nil {
		log.Fatalf("could not greet: %v", err)

	}
	if r.GetState() == true {
		log.Printf("Jugada enviada satisfactoriamente")
	}
	return true, err
}

func crearArchivo(path string) {
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
	fmt.Println("File Created Successfully", path)
}

func escribeArchivo(path string, texto string) {
	// Abre archivo usando permisos READ & WRITE
	var file, err = os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0660)
	if existeError(err) {
		return
	}
	defer file.Close()
	// Escribe algo de texto linea por linea
	_, err = file.WriteString(texto)
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

func (s *NameNodeServer) GuardarJugada(ctx context.Context, in *pb.JugadaToDataNode) (*pb.Response, error) {
	clientes := make([]pb.DataNameNodeServiceClient, 0)
	clientes = append(clientes, s.cliente1, s.cliente2, s.cliente3)
	pos := rand.Intn(len(clientes))
	cliente := clientes[pos]
	resp, err := EnviarJugadaADataNode(cliente, in, pos-1)
	// EnviarJugadaADataNode(cliente)
	log.Printf("Received: %v", in.GetId())
	respuesta := pb.Response{
		State: resp,
	}
	return &respuesta, err
}

func main() {
	// Set up a connection to the server.
	conn1, err := grpc.Dial(address1, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn1.Close()
	c1 := pb.NewDataNameNodeServiceClient(conn1)

	// Set up a connection to the server.
	conn2, err := grpc.Dial(address2, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn2.Close()
	c2 := pb.NewDataNameNodeServiceClient(conn2)

	// Set up a connection to the server.
	conn3, err := grpc.Dial(address3, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn3.Close()
	c3 := pb.NewDataNameNodeServiceClient(conn3)

	// clientes := make([]pb.DataNameNodeServiceClient, 0)
	// clientes = append(clientes, cliente1, cliente2, cliente3)

	// //Para distribuir a los clientes
	// cliente := clientes[rand.Intn(len(clientes))]

	// EnviarJugadaADataNode(cliente)
	// EnviarJugadaADataNode(cliente1)
	// EnviarJugadaADataNode(cliente2)
	// EnviarJugadaADataNode(cliente3)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLiderNameNodeServiceServer(s, &NameNodeServer{
		cliente1: c1,
		cliente2: c2,
		cliente3: c3,
	})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
