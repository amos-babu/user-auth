build:
	@go build -o bin/userauth

run: build
	@./bin/userauth

test:
	@go test -v ./...

clean:
	@rm -rf bin/