package main

import (
	"context"
	"log"
	"net"

	"github.com/krobus00/golang-grpc/pb"
	"google.golang.org/grpc"
)

type AuthService struct {
	pb.UnimplementedAuthServer
}

func (s *AuthService) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	return &pb.CreateUserResponse{
		User: &pb.User{
			Username: in.Username,
			FullName: in.FullName,
			Email:    in.Email,
		},
	}, nil
}

func (s *AuthService) LoginUser(ctx context.Context, in *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {

	return &pb.LoginUserResponse{
		User: &pb.User{
			Username: in.Username,
		},
		AccessToken: "TODO",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":17171")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAuthServer(s, &AuthService{})
	s.Serve(lis)
}
