package actions

import (
	. "github.com/StephanieSunshine/go-grpc-sendfile/api"
	"log"
	"context"
)

// Send generates response to a Ping request
func (s *Server) Send(ctx context.Context, in *FileRequestMessage) (*FileResponseMessage, error) {
        log.Println("Receive message: ", in)
  return &FileResponseMessage{}, nil
}

