package textprocessor

import (
	"context"
	"log"
)

type Server struct {
}

func (s *Server) GetTopTenUsedWords(ctx context.Context, message *Request) (*Response, error) {
	log.Printf("Received top ten used words: %s", message.body)
	return &Response{"Words"}, nil
}
