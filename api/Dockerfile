FROM golang:1.11.1

WORKDIR /app

COPY go.mod /app
RUN go mod download

COPY api.go /app
RUN go build

ENTRYPOINT ["./api"]
