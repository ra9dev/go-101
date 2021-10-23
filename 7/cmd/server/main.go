package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"lecture-7/api"
	"log"
	"net"
)

const (
	port = ":8080"
)

var (
	accounts = map[int64]*api.PersonalAccountResponse{
		1: {
			Id: 1,
			Name: "Sherkhan",
			Age: 23,
		},
		2: {
			Id: 2,
			Name: "Kirill",
			Age: 22,
		},
	}
)

type MailSender struct {
	api.UnimplementedMailServiceServer
}

func (ms *MailSender) MailSend(ctx context.Context, req *api.MailSendRequest) (*api.Empty, error) {
	log.Printf("Sending message to %s, text '%s'", req.To, req.Message)
	return &api.Empty{}, nil
}

type PersonalAccounter struct {
	api.UnimplementedPersonalAccountServiceServer
}

func (pa *PersonalAccounter) PersonalAccount(ctx context.Context, req *api.PersonalAccountRequest) (*api.PersonalAccountResponse, error) {
	if account, ok := accounts[req.Id]; ok {
		return account, nil
	}

	return nil, status.Errorf(codes.NotFound, fmt.Sprintf("account with id %d does not exist", req.Id))
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("cannot listen to %s: %v", port, err)
	}
	defer listener.Close()

	grpcServer := grpc.NewServer()
	mailSender := new(MailSender)
	personalAccounter := new(PersonalAccounter)

	api.RegisterMailServiceServer(grpcServer, mailSender)
	api.RegisterPersonalAccountServiceServer(grpcServer, personalAccounter)

	log.Printf("Serving on %v", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve on %v: %v", listener.Addr(), err)
	}
}
