PACKAGE := dayone

init:
	go mod init $(PACKAGE)

check-version:
	go --version

go-env:
	go env

run:
	go run main.go

mod-tidy:
	go mod tidy

mockgen:
	go get -u github.com/golang/mock/mockgen

gen mock:
	mockgen -source=./internal/handler/handler.go -destination=./internal/handler/mock_handler.go -package=handler