FROM golang:1.12-alpine
WORKDIR /go/src/app
ADD src src
CMD ["go", "run", "src/main/main.go"]