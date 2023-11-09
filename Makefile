build:
	@go build -o bin/gosandbox
run: build
	@./bin/gosandbox
