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
