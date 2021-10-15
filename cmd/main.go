package main

import (
	_ "go-microservice/docs"
	handler "go-microservice/pkg/handlers"
	"go-microservice/pkg/proto"
	"log"

	"google.golang.org/grpc"
)

// @title Telegram Bot task
// @version 1.0
// @description Send messages to telegram group or channel

// @host localhost:8081
// @BasePath /

func main() {
	messageServer, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := proto.NewMessageServiceClient(messageServer)
	handler.InitRoutes(client)
}
