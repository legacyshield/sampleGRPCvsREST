package main


import (
	"net"
	"fmt"
	"google.golang.org/grpc"
	pb "sampleGRPCvsREST/profile"

	"golang.org/x/net/context"
)


type server struct {}

func(s *server) CreateProfile(ctx context.Context, in *pb.ProfileRequest) (*pb.ProfileResponse, error){
		return &pb.ProfileResponse{Message: "Created profile with name " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Println("error in connection", err)
	}

	s := grpc.NewServer()
	pb.RegisterProfileServer(s, &server{})
	s.Serve(lis)
}