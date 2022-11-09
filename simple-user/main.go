package main
/*
#include "sum.h"
*/
import "C"

import (
	"context"
	"log"
	"net"
	"fmt"
	"google.golang.org/grpc"

	"github.com/samguya/grpc-example/data"
	userpb "github.com/samguya/grpc-example/protos/v1/user"
)

const portNumber = "9000"

type userServer struct {
	userpb.UserServer
}

// GetUser returns user message by user_id
func (s *userServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	userID := req.UserId
	var a, b int =  3, 2
	r := C.sum(C.int(a), C.int(b))
	fmt.Println(r)
	r = C.delfunc(C.int(a), C.int(b))
	fmt.Println(r)
	var userMessage *userpb.UserMessage
	for _, u := range data.Users {
		if u.UserId != userID {
			continue
		}
		userMessage = u
		break
	}

	return &userpb.GetUserResponse{
		UserMessage: userMessage,
	}, nil
}

// ListUsers returns all user messages
func (s *userServer) ListUsers(ctx context.Context, req *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	userMessages := make([]*userpb.UserMessage, len(data.Users))
	for i, u := range data.Users {
		userMessages[i] = u
	}

	return &userpb.ListUsersResponse{
		UserMessages: userMessages,
	}, nil
}

//export Init
func Init() {
        lis, err := net.Listen("tcp", ":"+portNumber)
        if err != nil {
                log.Fatalf("failed to listen: %v", err)
        }

        grpcServer := grpc.NewServer()
        userpb.RegisterUserServer(grpcServer, &userServer{})

        log.Printf("start gRPC server on %s port", portNumber)
        if err := grpcServer.Serve(lis); err != nil {
                log.Fatalf("failed to serve: %s", err)
        }
}

func main() {
}
