package main

import (
	"context"
	"fmt"
	student_service "gRPC_demo/grpc/idl/my_proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
)

func TestGrpc(t *testing.T) {
	conn, err := grpc.NewClient(
		"127.0.0.1:1234",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		t.Fatal(err)
	}
	client := student_service.NewStudentServiceClient(conn)
	info, err := client.GetStuInfo(context.Background(), &student_service.Request{StudentId: "1"})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(info)
}
