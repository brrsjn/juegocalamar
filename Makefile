gen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/*.proto

clean:
	rm pb/*.go
	rm nameNode/dataNode1/jugadas/*txt
	rm nameNode/dataNode2/jugadas/*txt
	rm nameNode/dataNode3/jugadas/*txt

run:
	go run main.go

server:
	go run service/lider_service.go

jugador_cliente:
	go run jugador/jugador.go --cli

jugador_bot:
	go run jugador/jugador.go

pozo:
	go run pozo/pozo.go

dataNode1:
	go run nameNode/dataNode1/dataNode1.go

dataNode2:
	go run nameNode/dataNode2/dataNode2.go

dataNode3:
	go run nameNode/dataNode3/dataNode3.go

name_node:
	go run nameNode/nameNode.go