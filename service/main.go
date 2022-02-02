package main

import (
	"log"
	"net"

	"github.com/aguleria91/mangtas-go-test/service/textprocessor"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on pont 9000: %v", err)
	}

	s := textprocessor.Server{}

	grpcServer := grpc.NewServer()

	textprocessor.RegisterTextProcessorServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve grpc server over 9000: %v", err)
	}

}

// // getTopTenUsedWords responds with a list of top ten used words as key and their number of occurrence as value
// func postTopTenUsedWords(c *gin.Context) {
// 	buf := new(bytes.Buffer)
// 	buf.ReadFrom(c.Request.Body)
// 	body := buf.String()

// 	// Remove punctuations from the paragraph
// 	replacer := strings.NewReplacer(",", "", ".", "", ";", "")
// 	body = replacer.Replace(body)

// 	words := strings.Fields(body)

// 	// Declare a map for words
// 	wordMap := make(map[string]int)

// 	// For loop to find dulpicate words as key and their occurrence as number.
// 	for _, word := range words {
// 		if v, ok := wordMap[word]; ok {
// 			wordMap[word] = v + 1
// 		} else {
// 			wordMap[word] = 1
// 		}
// 	}

// 	c.JSON(200, gin.H{"top_ten_words_used": wordMap})
// }
