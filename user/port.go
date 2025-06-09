//go:generate bash -c "mockgen -source=$(basename ${GOFILE} .go).go -package=$(go list -f '{{.Name}}') -destination=$(basename ${GOFILE} .go)_mock_test.go"
package user

import (
	"context"
	"dayone/gensql/postgresdb"
)

type getUsername interface {
	GetUsername(id string) string
}
type dbGetter interface {
	GetAuthor(ctx context.Context, id int64) (postgresdb.Author, error)
}
