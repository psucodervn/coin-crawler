FROM golang:1.11.1

WORKDIR /app

COPY go.mod /app
RUN go mod download

COPY worker.go /app
RUN go build

ENTRYPOINT ["./worker"]
