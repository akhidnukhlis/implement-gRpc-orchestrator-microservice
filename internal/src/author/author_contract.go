package author

import (
	"context"

	"github.com/akhidnukhlis/implement-gRpc-microservice-orchestrator/internal/entity/author_entity"
)

type Service interface {
	InsertNewAuthor(ctx context.Context, payload *author_entity.AuthorRequest) error
	FindAuthor(ctx context.Context, authorID string) (*author_entity.Authors, error)
}
