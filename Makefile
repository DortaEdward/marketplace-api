build:
	@go build -o bin/src

run: build
	@bin/src
