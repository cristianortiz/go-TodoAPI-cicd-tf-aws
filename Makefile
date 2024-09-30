build:
	CGO_ENABLED=0 env GOOS=linux go build -o bin/todoAPI
#builnd an run the app locally
run: build
	./bin/todoAPI
#run mongoDB with docker
mongo:
	docker compose up --build -d
#stop mongo container
stop:
	docker compose down
test:
	go test -v ./...
tcover:
	go test  -coverprofile=c.out | go tool cover -html=c.out

    