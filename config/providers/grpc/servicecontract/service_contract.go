package servicecontract

import (
	"github.com/akhidnukhlis/implement-gRpc-proto-bank/grpc/pb"
)

type GrpcContract struct {
	Author pb.AuthorServiceClient
}

func NewGrpcService(author pb.AuthorServiceClient) *GrpcContract {
	return &GrpcContract{
		Author: author,
	}
}
