gen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/*.proto
clean:
	rm pb/*.go

run:
	go run main.go

server:
	go run service/lider_service.go

jugador_cliente:
	go run jugador/jugador.go --cli

jugador_bot:
	go run jugador/jugador.go