package main

import (
	"fmt"
	"log"
	"os"

	amqp "github.com/streadway/amqp"
)

func existeError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
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

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var path = "client/pozo.txt"
			var file, err = os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0660)
			var escr string
			escr = string(d.Body)

			if existeError(err) {
				return
			}
			defer file.Close()
			// Escribe algo de texto linea por linea
			_, err = file.WriteString(escr + "\n")
			if existeError(err) {
				return
			}

			// Salva los cambios
			err = file.Sync()
			if existeError(err) {
				return
			}
			fmt.Println("Archivo actualizado existosamente.")
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
