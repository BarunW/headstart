build:
	go build -o bin/headstart .
run:
	go run .
test:
	go test -v .

clean:
	rm -rf bin
