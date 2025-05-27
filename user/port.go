//go:generate bash -c "mockgen -source=$(basename ${GOFILE} .go).go -package=$(go list -f '{{.Name}}') -destination=$(basename ${GOFILE} .go)_mock_test.go"
package user

type getUsername interface {
	GetUsername(id string) string
}
