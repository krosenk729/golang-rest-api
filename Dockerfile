FROM golang:alpine
ADD . /go/src/app
WORKDIR /go/src/app
ENV PORT=5050
CMD ["go", "run", "main.go"]