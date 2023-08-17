package tests

import (
	"context"
	"github.com/akhidnukhlis/implement-gRpc-microservice-orchestrator/config/providers/grpc/servicecontract"
	"github.com/akhidnukhlis/implement-gRpc-microservice-orchestrator/helpers/errorcodehandling"
	"github.com/akhidnukhlis/implement-gRpc-microservice-orchestrator/internal/entity"
	"github.com/jinzhu/gorm"
)

type Fields struct {
	Repo *servicecontract.GrpcContract
	Err  *errorcodehandling.CodeError
}

type Args struct {
	Ctx     context.Context
	payload *entity.AuthorRequest
	DB      *gorm.DB
}
