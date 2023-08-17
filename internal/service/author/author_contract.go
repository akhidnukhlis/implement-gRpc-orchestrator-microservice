package author

import (
	"context"
	"github.com/akhidnukhlis/implement-gRpc-orchestrator-microservice/internal/entity"
)

type Service interface {
	CreateNewAuthor(ctx context.Context, payload *entity.AuthorRequest) error
	FindAuthor(ctx context.Context, authorID string) (*entity.Authors, error)
}
