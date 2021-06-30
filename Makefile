all: main

run: main
	./bin/main

main: fmt vet tidy
	GOPATH=/Users/D072532/Documents/4Gewinnt
	go build -o bin ./...

fmt:
	go fmt ./... 

vet:
	go fmt ./...

tidy:
	go mod tidy

test: 
	go test ./...
