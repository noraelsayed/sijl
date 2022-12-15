FROM golang:1.19

RUN mkdir -p /usr/src/sijl
COPY . /usr/src/sijl
WORKDIR /usr/src/sijl
RUN go mod tidy
ENTRYPOINT go run cmd/sijl.go
