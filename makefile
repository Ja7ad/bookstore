.PHONY: fmt test vet lint build clean


fmt:
	go fmt ./...

test:
	go test ./... -v

vet:
	go vet ./...

build:
	go build -o build/bookstore cmd/bookstore/main.go
	go build -o build/oauth cmd/oauth/main.go
	cp -r ./README.md build/release

clean: fmt vet build