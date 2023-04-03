package app

import (
	"github.com/akhidnukhlis/implement-gRpc-microservice-orchestrator/internal/entity/author_entity"
)

// SetMigrationTable is used to register entity model which want to be migrate
func SetMigrationTable() []interface{} {
	var migrationData = []interface{}{
		&author_entity.Author{},
	}

	return migrationData
}
