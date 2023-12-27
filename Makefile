
APP_NAME=rinhabackend2023

build:
	go test ./...
	CGO_ENABLED=0 GOOS=linux go build -o bin/$(APP_NAME)

swagger:
	swag init -g internal/http/api.go
	swag fmt