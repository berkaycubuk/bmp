build:
	@go build -o bin/bmp main.go

run: build
	@./bin/bmp
