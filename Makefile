
APP_NAME=rinhabackend2023

build:
	go clean
	CGO_ENABLED=0 GOOS=linux go build -o bin/$(APP_NAME) cmd/rinhabackend/main.go

swagger:
	swag init -g internal/http/api.go
	swag fmt

docker-build:
	docker build -t marcusadriano/rinhabackend-go:latest .