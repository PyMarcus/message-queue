run: build
	@./bin/message_queue

build:
	@go build -o bin/message_queue
