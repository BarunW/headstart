build:
	go build -o bin/headstart .
run:
	go run .
test:
	go test -v .

clean_bin:
	rm -rf bin
clean_cache:
	go clean -testcache
	go clean -cache
