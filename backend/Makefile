include .env
export $(shell sed 's/=.*//' .env)
build:
	@go build -o bin/nft-ticket-verification ./cmd

run: build
	@./bin/nft-ticket-verification

test:
	@go test -v ./...
