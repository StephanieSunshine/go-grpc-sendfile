package api

import (
  "log"
  "golang.org/x/net/context"
)

// Server represents the gRPC server
type Server struct {
}

// SayHello generates response to a Ping request
func (s *Server) Send(ctx context.Context, in *FileRequestMessage) (*FileResponseMessage, error) {
  log.Printf("Receive message %s", in.Greeting)
  return &FileResponseMessage{}, nil
}
