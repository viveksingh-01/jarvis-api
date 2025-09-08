#build
build:
	@go build -o bin/jarvis-api cmd/main.go

#run
run: build
	@./bin/jarvis-api