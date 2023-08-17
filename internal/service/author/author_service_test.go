package author

import (
	"github.com/akhidnukhlis/implement-gRpc-orchestrator-microservice/external/seeders"
	"github.com/akhidnukhlis/implement-gRpc-orchestrator-microservice/external/tests"
	"github.com/akhidnukhlis/implement-gRpc-orchestrator-microservice/internal/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_service_CreateNewAuthor(t *testing.T) {

	tt := struct {
		fields tests.Fields
		args   tests.Args
	}{}

	seedAuthor, _ := seeders.SeedAuthor(tt.args.DB)

	t.Run("if repo was successfully recorded, it should not return error", func(t *testing.T) {
		s := &service{
			repo: tt.fields.Repo,
			err:  tt.fields.Err,
		}

		err := s.CreateNewAuthor(tt.args.Ctx, &entity.AuthorRequest{
			Name:     seedAuthor.Name,
			Nickname: seedAuthor.Nickname,
			Email:    seedAuthor.Email,
		})
		assert.NoError(t, err)
	})
}

func Test_service_FindAuthor(t *testing.T) {
	tt := struct {
		fields tests.Fields
		args   tests.Args
	}{}

	seedAuthor, _ := seeders.SeedAuthor(tt.args.DB)

	t.Run("if repo was exists, it should not return error", func(t *testing.T) {
		s := &service{
			repo: tt.fields.Repo,
			err:  tt.fields.Err,
		}

		result, err := s.FindAuthor(tt.args.Ctx, seedAuthor.ID)
		assert.NoError(t, err)
		assert.Equal(t, seedAuthor.ID, result.ID)
	})
}
