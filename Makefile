generate:
	go generate ./...
tools:
	go get github.com/securego/gosec/v2/cmd/gosec
	go get github.com/fzipp/gocyclo/cmd/gocyclo
	go get golang.org/x/lint/golint
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.32.2
fmt:
	go fmt ./...

lint: generate fmt
	go vet ./...
	gocyclo -over 20 -ignore ".*_test.go" .
	golint ./...
	gosec -severity medium -quiet ./...

test: lint
	go test ./...
.PHONY: generate fmt test