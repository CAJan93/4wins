all: main

run: main
	./bin/fourwins

# Build manager binary
main: # fmt vet tidy
	GOPATH=/Users/D072532/Documents/4Gewinnt
	go build -o bin/fourwins src/main/main.go && go version -m bin/fourwins

# TODO: as a workaround I named all packages for fmt and vet as ./... will fail in the pipeline due to the vendor folder
# Run go fmt against code
fmt:
	go fmt ./... 

# Run go vet against code
vet:
	go fmt ./...

tidy:
	go mod tidy
