package main

import (
	"context"
	"log"
	"net"

	pb "github.com/vinusstar/grpcrest/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8888"
)

type aliveService struct{}

func (s *aliveService) GetStatus(ctx context.Context, in *pb.Empty) (*pb.AliveResponse, error) {
	return &pb.AliveResponse{Status: true}, nil
}

type userService struct{}

func (s *userService) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{
		Id:   in.Id,
		Name: "Alice",
		Age:  20,
	}, nil
}

func (s *userService) GetUserByGroup(ctx context.Context, in *pb.UserGroupRequest) (*pb.UsersResponse, error) {
	return &pb.UsersResponse{
		Group: in.Group,
		Users: []*pb.UserResponse{
			{Name: "Alice", Age: 20},
			{Name: "Bob", Age: 24},
		},
	}, nil
}

func (s *userService) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.Empty, error) {
	log.Printf("update body is {id: %s, name: %s, age: %d}\n", in.Id, in.Name, in.Age)
	return &pb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterAliveServiceServer(s, new(aliveService))
	pb.RegisterUserServiceServer(s, new(userService))
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}

}
