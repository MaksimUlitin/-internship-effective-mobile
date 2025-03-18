.PHONY: run swag-generate all build docker-build docker-up docker-down clean

all: swag-generate run

run:
	go run cmd/main.go

swag-generate:
	cd cmd && swag init -g ../cmd/main.go -d ../config,../internal/models,../internal/controllers,../internal/storage/database -o ../docs

build:
	go build -o music-library-app ./cmd/main.go

docker-build:
	docker-compose build

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

docker-logs:
	docker-compose logs -f

clean:
	rm -f music-library-app
	docker-compose down -v

docker-all: docker-build docker-up