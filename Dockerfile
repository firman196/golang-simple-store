FROM golang:1.20.3-alpine3.17
WORKDIR /app
ENV GOPATH=/app
COPY . .

RUN go build -o main main.go