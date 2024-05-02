#oficial go image
FROM golang:1.22:bookworm
#set working directory inside container
WORKDIR /app
#copy go mod and go sum
COPY go.mod .
COPY go.sum .

#downlod depedencies
RUN go mod downlod

#copy the source code into container
COPY . .
#buil the go app
RUN go build -o main .
#command to run the exe
CMD ["./main"]
