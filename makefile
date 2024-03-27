dev:
	go run main.go

test:
	go test -v -race ./...
	go vet ./...

generate:
	go mod tidy
	go generate ./...
	swag init

build:
	docker build -t wallet .

run:
	docker run -p 8080:8080 wallet
