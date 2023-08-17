package client

import (
	"github.com/akhidnukhlis/implement-gRpc-proto-bank/grpc/pb"
	"log"
	"os"

	"google.golang.org/grpc"
)

func ServiceUser() pb.AuthorServiceClient {
	port := os.Getenv("GRPC_CLIENT_AUTHOR")
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return pb.NewAuthorServiceClient(conn)
}
