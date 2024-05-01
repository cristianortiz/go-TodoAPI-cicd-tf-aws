build:
	go build -o bin/todoAPI
run: build
	./bin/todoAPI