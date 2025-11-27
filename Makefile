build:
	@go build -o bin/userauth ./cmd

run: build
	@./bin/userauth

test:
	@go test -v ./...

clean:
	@rm -rf bin/