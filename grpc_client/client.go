package main

import (
	"google.golang.org/grpc"
	"fmt"
	pb "sampleGRPCvsREST/profile"
	"golang.org/x/net/context"
)


func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Error in creating grpc connection")
	}

	defer conn.Close()

	c := pb.NewProfileClient(conn)

	r, err := c.CreateProfile(context.Background(), &pb.ProfileRequest{Name: "test", Age: 21, City: "testcity", Country: "testcountry"})
	if err != nil {
		fmt.Println("Error in serving grpc call")
	}

	fmt.Println("Message", r.Message)
}