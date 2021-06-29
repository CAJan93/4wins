all: main

run: main
	./bin/fourwins

main: fmt vet tidy
	GOPATH=/Users/D072532/Documents/4Gewinnt
	go build -o bin/fourwins ./... && go version -m bin/fourwins

fmt:
	go fmt ./... 

vet:
	go fmt ./...

tidy:
	go mod tidy

test: 
	go test ./...
