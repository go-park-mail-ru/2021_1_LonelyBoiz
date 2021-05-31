FROM golang:latest
RUN mkdir app
ADD . ./app
WORKDIR ./app
RUN go mod tidy

ENTRYPOINT go run cmd/auth_server/main.go

EXPOSE 5400