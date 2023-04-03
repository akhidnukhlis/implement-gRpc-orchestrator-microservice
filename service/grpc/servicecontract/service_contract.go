package servicecontract

import (
	"github.com/akhidnukhlis/implement-gRpc-microservice/grpc/pb"
)

type GrpcContract struct {
	Author pb.AuthorServiceClient
}

func NewGrpcService(author pb.AuthorServiceClient) *GrpcContract {
	return &GrpcContract{
		Author: author,
	}
}
