setup:
	cp .env.example .env

lint:
	go fmt ./...

build:
	go build -v .

run-docker:
	docker-compose build && docker-compose up
