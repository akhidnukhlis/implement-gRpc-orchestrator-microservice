package author

import (
	"context"

	"github.com/akhidnukhlis/implement-gRpc-microservice-orchestrator/helpers/errorcodehandling"
	"github.com/akhidnukhlis/implement-gRpc-microservice-orchestrator/helpers/unique"
	"github.com/akhidnukhlis/implement-gRpc-microservice-orchestrator/internal/entity"
	"github.com/akhidnukhlis/implement-gRpc-microservice-orchestrator/internal/entity/author_entity"
	"github.com/akhidnukhlis/implement-gRpc-microservice-orchestrator/service/grpc/servicecontract"
	"github.com/akhidnukhlis/implement-gRpc-microservice/grpc/pb"
)

type service struct {
	author *servicecontract.GrpcContract
	err    *errorcodehandling.CodeError
}

func NewService(author *servicecontract.GrpcContract) *service {
	return &service{
		author: author,
	}
}

// InsertNewAuthor represents algorithm to register new author
func (s *service) InsertNewAuthor(ctx context.Context, payload *author_entity.AuthorRequest) error {

	err := author_entity.AuthorRequestValidate(payload)
	if err != nil {
		return err
	}

	author := &pb.CreateAuthorRequest{
		Name:     payload.Name,
		Nickname: payload.Nickname,
		Email:    payload.Email,
	}

	_, err = s.author.Author.ServiceRegisterAuthor(ctx, author)
	if err != nil {
		return err
	}

	return nil
}

// FindAuthor represents algorithm to find author by id
func (s *service) FindAuthor(ctx context.Context, authorID string) (*author_entity.Authors, error) {
	if err := unique.ValidateUUID(authorID); err != nil {
		return nil, entity.ErrAuthorNotExist
	}

	payload := &pb.FindAuthorByIdRequest{
		Id: authorID,
	}

	author, err := s.author.Author.ServiceFindAuthorById(ctx, payload)
	if err != nil {
		return nil, err
	}

	usrData := &author_entity.Authors{
		ID:        authorID,
		Name:      author.Data.Name,
		Nickname:  author.Data.Nickname,
		Email:     author.Data.Email,
		CreatedAt: author.Data.CreatedAt,
		UpdatedAt: author.Data.UpdatedAt,
	}

	return usrData, nil
}
