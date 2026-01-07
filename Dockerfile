FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum .
RUN go mod download
COPY main.go .

RUN go build -v -o logging-mock-service .

ENTRYPOINT [ "/app/logging-mock-service" ]
CMD ["--delay", "200", "--output", "20000", "--verbose"]

