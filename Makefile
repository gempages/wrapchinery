upgrade:
	go get -u github.com/gempages/go-helper@production
	go mod tidy

gotest:
	MODE=test go test -cover -race ./...
