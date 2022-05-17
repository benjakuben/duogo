OUTPUT = ./main

.PHONY: clean
clean:
	rm -f $(OUTPUT)

.PHONY: install
install:
	go get ./...

graphql: ./main.go
	go build -o $(OUTPUT) ./main.go

# compile the code to run in Lambda (local or real)

.PHONY: lambda
lambda:
	GOOS=linux GOARCH=amd64 $(MAKE) graphql

.PHONY: build
build: clean lambda

.PHONY: local
local: build
	sam local start-api

.PHONY: migrate-up
migrate-up:
	migrate -path db/migrations -database "postgres://postgres@localhost:5432/strutdb?sslmode=disable" up

.PHONY: migrate-down
migrate-down:
	migrate -path db/migrations -database "postgres://postgres@localhost:5432/strutdb?sslmode=disable" down