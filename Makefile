build:
	go build -o bin/todoAPI
#builnd an run the app locally
run-go: build
	./bin/todoAPI
#run the APi with docker
run:
	docker compose up -d
#stop the app stopping docker
stop:
	docker compose down