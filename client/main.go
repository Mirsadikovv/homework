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

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "file/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// const (
// 	defaultName       = "world"
// 	default_a   int64 = 123
// 	default_b   int64 = 121
// )

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewFileManagerClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// r, err := c.CreateFile(ctx, &pb.CreateRequest{Path: "/home/mirodil/Desktop/src/github/month4/homeworks/lesson1", FileName: "new1.txt"})
	// if err != nil {
	// 	log.Fatalf("response: %v", err)
	// }

	// r, err := c.IsFileExists(ctx, &pb.CheckRequest{Path: "/home/mirodil/Desktop/src/github/month4/homeworks/lesson1", FileName: "new1.txt"})
	// if err != nil {
	// 	log.Fatalf("could not check: %v", err)
	// }
	// log.Printf("Response: %v", r.GetMessage())

	// r, err := c.WriteToFile(ctx, &pb.WriteRequest{Path: "/home/mirodil/Desktop/src/github/month4/homeworks/lesson1", FileName: "new1.txt", Data: "asdsad"})
	// if err != nil {
	// 	log.Fatalf("could not write: %v", err)
	// }
	// log.Printf("Response: %v", r.GetMessage())

	r, err := c.ReadFromFile(ctx, &pb.ReadRequest{Path: "/home/mirodil/Desktop/src/github/month4/homeworks/lesson1", FileName: "new1.txt"})
	if err != nil {
		log.Fatalf("could not read: %v", err)
	}
	log.Printf("Response: %v", r.GetMessage())
}
