package web

import (
	"fmt"
	"log"
	"net"
	"net/http"

	auth_controller "github.com/simple-distributed-project/web/auth/controller"
	"github.com/simple-distributed-project/web/auth/gapi"
	"github.com/simple-distributed-project/web/auth/pb"
	userrepo "github.com/simple-distributed-project/web/auth/repository"
	bcrypt "golang.org/x/crypto/bcrypt"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartService() {
	http.HandleFunc("/", auth_controller.Home)
	http.HandleFunc("/login", auth_controller.Login)
	http.HandleFunc("/register", auth_controller.Register)
	http.HandleFunc("/profile", auth_controller.Profile)
	http.HandleFunc("/follow", auth_controller.Follow)
	http.HandleFunc("/unfollow", auth_controller.Unfollow)
	http.HandleFunc("/tweet", auth_controller.Tweet)
	http.HandleFunc("/feed", auth_controller.Feed)
	http.HandleFunc("/logout", auth_controller.Logout)
	password := "aa"
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	if err != nil {
		fmt.Println("Error While Hashing Password, Try Again")
	}
	password = string(hash)

	err = userrepo.CreateUser("virenmparmar@gmail.com", password)
	err = userrepo.CreateUser("vmp2018@nyu.edu", password)
	err = http.ListenAndServe(":8000", nil) // set listen port
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	} else {
		fmt.Println("Server starts at: localhost:8000")
	}
}

func StartGRPCServer() {
	server, err := gapi.NewServer()
	if err != nil {
		log.Fatalf("cannot start server: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterProjectServerServer(grpcServer, server)
	reflection.Register(grpcServer)
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}

}
