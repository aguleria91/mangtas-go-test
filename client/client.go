package main

import (
	"context"
	"log"

	"github.com/aguleria91/mangtas-go-test/service/textprocessor"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure)

	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}

	defer conn.Close()

	c := textprocessor.NewTextProcessorServiceClient(conn)

	message := textprocessor.Message{}

	response, err := c.GetTopTenUsedWords(context.Background(), &message)

	if err != nil {
		log.Fatalf("Error when calling GetTopTenUsedWords %s", err)
	}

	log.Printf("Response from Server: %s", response.Body)
}
