FROM golang:1.11.1

RUN go get github.com/pilu/fresh

WORKDIR /app

ENV RUNNER_TMP_PATH=/tmp
CMD ["fresh"]
