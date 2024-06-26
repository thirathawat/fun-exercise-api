dev:
	go run main.go

test:
	go vet ./...
	go test -v -race ./...

generate:
	go mod tidy
	go generate ./...
	swag init

build:
	docker build -t wallet .

run:
	docker run -p 8080:8080 wallet
