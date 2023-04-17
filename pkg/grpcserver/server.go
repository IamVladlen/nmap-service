package grpcserver

import (
	"net"

	"github.com/IamVladlen/nmap-service/pkg/logger"
	"google.golang.org/grpc"
)

type Server struct {
	*grpc.Server

	notify chan error
}

// New creates and starts gRPC server.
func New(port string, log *logger.Log) (*Server, error) {
	s := &Server{
		grpc.NewServer(),
		make(chan error, 1),
	}

	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return &Server{}, err
	}

	go func() {
		s.notify <- s.Serve(l)
		close(s.notify)
	}()

	return s, nil
}

// Notify returns an error that occured 
// in the gRPC server.
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Stop gracefully stops the gRPC server.
func (s *Server) Stop() {
	s.GracefulStop()
}
