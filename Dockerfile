# Dockerfile References: https://docs.docker.com/engine/reference/builder/

FROM golang:alpine
ADD . /go/src/app
WORKDIR /go/src/app

RUN go build

CMD ["go", "run", "main.go"]