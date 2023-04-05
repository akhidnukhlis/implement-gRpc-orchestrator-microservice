package author

import (
	"context"

	"github.com/akhidnukhlis/implement-gRpc-microservice-orchestrator/helpers/errorcodehandling"
	"github.com/akhidnukhlis/implement-gRpc-microservice-orchestrator/helpers/unique"
	"github.com/akhidnukhlis/implement-gRpc-microservice-orchestrator/internal/entity"
	"github.com/akhidnukhlis/implement-gRpc-microservice-orchestrator/service/grpc/servicecontract"
	"github.com/akhidnukhlis/implement-gRpc-microservice/grpc/pb"
)

type service struct {
	repo *servicecontract.GrpcContract
	err  *errorcodehandling.CodeError
}

func NewService(author *servicecontract.GrpcContract) *service {
	return &service{
		repo: author,
	}
}

// CreateNewAuthor represents algorithm to register new repo
func (s *service) CreateNewAuthor(ctx context.Context, payload *entity.AuthorRequest) error {

	err := entity.AuthorRequestValidate(payload)
	if err != nil {
		return err
	}

	author := &pb.CreateAuthorRequest{
		Name:     payload.Name,
		Nickname: payload.Nickname,
		Email:    payload.Email,
	}

	_, err = s.repo.Author.ServiceRegisterAuthor(ctx, author)
	if err != nil {
		return err
	}

	return nil
}

// FindAuthor represents algorithm to find repo by id
func (s *service) FindAuthor(ctx context.Context, authorID string) (*entity.Authors, error) {
	if err := unique.ValidateUUID(authorID); err != nil {
		return nil, entity.ErrAuthorNotExist
	}

	payload := &pb.FindAuthorByIdRequest{
		Id: authorID,
	}

	author, err := s.repo.Author.ServiceFindAuthorById(ctx, payload)
	if err != nil {
		return nil, err
	}

	usrData := &entity.Authors{
		ID:        authorID,
		Name:      author.Data.Name,
		Nickname:  author.Data.Nickname,
		Email:     author.Data.Email,
		CreatedAt: author.Data.CreatedAt,
		UpdatedAt: author.Data.UpdatedAt,
	}

	return usrData, nil
}
