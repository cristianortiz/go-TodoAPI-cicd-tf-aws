#build stage
FROM golang:1.22-alpine3.20 AS builder
#upx cto comprise image size
RUN apk add --no-cache git
#set working directory inside container
WORKDIR /app
#copy go mod and go sum
COPY go.mod  go.sum  ./

#downlod depedencies
RUN go mod download
# Copy the source code. Note the slash at the end, as explained in
COPY  . ./

# Copy the db/migrations directory
#COPY database/migrations database/migrations

#buil the go app
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o todo-api

#RUN upx /todo-api

#final stage
FROM alpine:3.20
RUN apk update --no-cache add ca-certificates
RUN apk --no-cache add ca-certificates
WORKDIR /app

COPY --from=builder /app .
# Run
ENTRYPOINT ["./todo-api"]