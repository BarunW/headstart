FILE_NAME ?= main.

#target command
all : test

run:
	@go run .

test:
	@go test -v $(FILE_NAME)

build:
	@go build .

