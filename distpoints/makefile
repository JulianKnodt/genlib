exec: build
	./distpoints

build: main.go fmt
	go build

fmt:
	go fmt ./...
	go vet ./...
