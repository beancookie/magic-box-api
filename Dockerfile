FROM golang:alpine AS builder

ARG VERSION=dev

WORKDIR /app

ADD go.mod .
ADD go.sum .

# RUN go env -w  GOPROXY=https://goproxy.cn,direct
RUN go mod download
ADD . .

ENV GO111MODULE=on
RUN GOOS=linux GOARCH=amd64 go build -o main -ldflags=-X=main.version=${VERSION} main.go


FROM alpine:latest
WORKDIR /root
COPY --from=builder /app/main .
CMD ["./main"]