package gapi

import (
	"context"
	"fmt"
	"strings"

	"github.com/simple-distributed-project/web/auth/pb"
	userrepo "github.com/simple-distributed-project/web/auth/repository"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) RegisterService(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	username := req.GetUsername()
	password := req.GetPassword()

	user, err := userrepo.FindUser(strings.TrimSpace(username))
	if user.Username == username {
		return nil, status.Errorf(codes.Internal, "User already exists: %s", err)
	}
	if err.Error() != "User does not exist" {
		return nil, status.Errorf(codes.Internal, "user does not exist: %s", err)
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate hash of password: %s", err)
	}
	password = string(hash)

	err = userrepo.CreateUser(strings.TrimSpace(username), password)
	if err != nil {
		fmt.Println(err)
		return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
	}

	rsp := &pb.RegisterResponse{
		Message: "User created successfully",
	}
	return rsp, nil
}
