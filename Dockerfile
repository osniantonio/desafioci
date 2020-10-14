FROM golang:1.12-alpine as builder
ADD src/app /go/src/app
WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w"

FROM scratch
ADD src/app /go/src/app
WORKDIR /go/src/app
COPY --from=builder /go/src/app .
CMD ["./app"]