package main

import (
	"log"
	"testing"
	pb "sampleGRPCvsREST/profile"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"fmt"
	"net/http"
	"crypto/tls"
	"bytes"
)

func BenchmarkGRPCCreateProfile(b *testing.B) {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Error in creating grpc connection")
	}

	defer conn.Close()
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewProfileClient(conn)

	// run grpc calls against it
	for i := 0; i < b.N; i++ {
		client.CreateProfile(context.Background(), &pb.ProfileRequest{Name: "test", Age: 21, City: "testcity", Country: "testcountry"})
	}

}

func BenchmarkRESTCreateProfile(b *testing.B) {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	buf := bytes.NewBufferString(`
		{
			"name":"test",
			"age":1,
			"height":1
		}
	`)
	// run http posts against it
	for i := 0; i < b.N; i++ {
		client.Post("https://localhost:3000/profile", "application/json", buf)
	}
}