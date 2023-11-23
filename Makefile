hello:
	echo "Hello, World"

run:
	go run cmd/app/main.go

build: 
	go build -o ./cmd/bin/silliChat cmd/app/main.go

test:

