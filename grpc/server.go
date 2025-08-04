package main

import (
	"context"
	"errors"
	"fmt"
	student_service "gRPC_demo/grpc/idl/my_proto"
	"google.golang.org/grpc"
	"net"
)

type StudentServer struct {
	student_service.UnimplementedStudentServiceServer
}

func (StudentServer) GetStuInfo(ctx context.Context, in *student_service.Request) (*student_service.Student, error) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("%v", err)
		}
	}()
	if len(in.StudentId) == 0 {
		return nil, errors.New("id is null")
	}
	return &student_service.Student{
		Name:   "zhangsan",
		Age:    28,
		Height: 1.75,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	student_service.RegisterStudentServiceServer(server, &StudentServer{})
	err = server.Serve(lis)
	if err != nil {
		panic(err)
	}
}
