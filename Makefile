
APP_NAME=rinhabackend2023
CACHE_SERVER=cacheserver

build:
	go clean
	CGO_ENABLED=0 GOOS=linux go build -o bin/$(APP_NAME) cmd/rinhabackend/main.go

build-cache:
	go clean
	CGO_ENABLED=0 GOOS=linux go build -o bin/$(CACHE_SERVER) cmd/cacheserver/main.go

swagger:
	swag init -g internal/http/api.go
	swag fmt

docker-build:
	docker build -t marcusadriano/rinhabackend-go:latest .
	docker build -t marcusadriano/cacheserver:latest -f Dockerfile.cache .
	docker scout cves local://marcusadriano/rinhabackend-go:latest