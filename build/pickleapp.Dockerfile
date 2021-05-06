FROM golang:latest
RUN mkdir app
ADD . ./app
WORKDIR ./app
RUN go mod tidy

ENTRYPOINT go run cmd/pickleapp/main.go

EXPOSE 8000