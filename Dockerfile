#oficial go image
FROM golang:1.22
#set working directory inside container
WORKDIR /app
#copy go mod and go sum
COPY go.mod  go.sum ./

#downlod depedencies
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY *.go ./
#buil the go app
RUN CGO_ENABLED=0 GOOS=linux go build -o /todo-api

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 8080

# Run
CMD ["/todo-api"]
