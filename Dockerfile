FROM golang:1.12-alpine as builder
ADD src/app /src/app
WORKDIR /src/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w"

FROM scratch
ADD src/app /src/app
WORKDIR /src/app
COPY --from=builder /src/app .
CMD ["./app"]