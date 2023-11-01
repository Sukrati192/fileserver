FROM golang:1.20-alpine AS builder
WORKDIR /fileserver
ENV GOPATH /fileserver/go/
ENV GOBIN /fileserver/go/bin
COPY . .
WORKDIR /fileserver/http
RUN CGO_ENABLED=0 GOOS=linux go build -o /fileserver/main -mod vendor

WORKDIR /fileserver/
EXPOSE 8000 
CMD ["/fileserver/main", "--port", "8000", "--config", "/fileserver/configmanager/configs/config.json"]

