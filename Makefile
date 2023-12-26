build_directory := bin

.PHONY: buildexec

#make bin dir and build the exec inside it
buildexec: setup build 

setup: 
	mkdir -p bin

build: setup
	go build -o bin/headstart .

.PHONY: run

#run
run:
	go build run .

#test
test:
	go test .
