package gapi

import "github.com/simple-distributed-project/web/auth/pb"

type Server struct {
	pb.UnimplementedProjectServerServer
}

// New Server creates a new gRPC server
func NewServer() (*Server, error) {
	server := &Server{}
	return server, nil
}
