/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	pb "file/proto"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedFileManagerServer
}

// func (s *server) SumOf(ctx context.Context, in *pb.SumRequest) (*pb.SumReply, error) {
// 	log.Printf("Received: %v, %v", in.GetA(), in.GetB())
// 	return &pb.SumReply{Message: in.GetA() + in.GetB()}, nil
// }

func (s *server) CreateFile(ctx context.Context, in *pb.CreateRequest) (*pb.CreateReply, error) {
	a, b := in.GetPath(), in.GetFileName()
	fullpath := a + "/" + b
	log.Printf("Received: %v, %v", a, b)
	content := fmt.Sprintf("Path -> %v", fullpath)

	file, err := os.Create(fullpath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return &pb.CreateReply{Message: content}, nil
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func (s *server) IsFileExists(ctx context.Context, in *pb.CheckRequest) (*pb.CheckReply, error) {
	a, b := in.GetPath(), in.GetFileName()
	fullpath := a + "/" + b
	log.Printf("Received: %v, %v", a, b)

	if FileExists(fullpath) {
		res := true
		return &pb.CheckReply{Message: res}, nil
	}
	res := false
	return &pb.CheckReply{Message: res}, nil
}

func (s *server) WriteToFile(ctx context.Context, in *pb.WriteRequest) (*pb.WriteReply, error) {
	a, b, c := in.GetPath(), in.GetFileName(), in.GetData()
	fullpath := a + "/" + b
	log.Printf("Received: %v, %v", a, b)

	file, err := os.OpenFile(fullpath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	_, err = file.WriteString(c)
	if err != nil {
		return nil, err
	}
	return &pb.WriteReply{Message: "written"}, nil
}

func (s *server) ReadFromFile(ctx context.Context, in *pb.ReadRequest) (*pb.ReadReply, error) {
	a, b := in.GetPath(), in.GetFileName()
	fullpath := a + "/" + b
	log.Printf("Received: %v, %v", a, b)

	file, err := os.Open(fullpath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return &pb.ReadReply{Message: string(data)}, nil
}
func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterFileManagerServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
