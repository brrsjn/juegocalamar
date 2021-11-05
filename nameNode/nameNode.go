package main

import (
	"context"
	"fmt"
	"juegocalamar/pb"
	"log"
	"math/rand"
	"os"
	"time"

	"google.golang.org/grpc"
)

const (
	address1 = "localhost:50061"
	address2 = "localhost:50062"
	address3 = "localhost:50063"
)

func EnviarJugadaADataNode(cliente pb.DataNameNodeServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := cliente.RegistraJugada(ctx, &pb.JugadaToDataNode{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)

	}
	if r.GetState() == true {
		log.Printf("Jugada enviada satisfactoriamente")
	}
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
	// Set up a connection to the server.
	conn1, err := grpc.Dial(address1, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn1.Close()
	cliente1 := pb.NewDataNameNodeServiceClient(conn1)

	// Set up a connection to the server.
	conn2, err := grpc.Dial(address2, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn2.Close()
	cliente2 := pb.NewDataNameNodeServiceClient(conn2)

	// Set up a connection to the server.
	conn3, err := grpc.Dial(address3, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn3.Close()
	cliente3 := pb.NewDataNameNodeServiceClient(conn3)

	clientes := make([]pb.DataNameNodeServiceClient, 0)
	clientes = append(clientes, cliente1, cliente2, cliente3)

	//Para distribuir a los clientes
	cliente = clientes[rand.Intn(len(clientes))]

	EnviarJugadaADataNode(cliente1)
	EnviarJugadaADataNode(cliente2)
	EnviarJugadaADataNode(cliente3)

}
