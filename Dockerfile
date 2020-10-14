FROM golang:1.12-alpine as builder
WORKDIR /workspace/gopath/src/app
ADD src/app /workspace/gopath/src/app
COPY . .
# para a imagem reduzida
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w"

FROM scratch
WORKDIR /workspace/gopath/src/app
ADD src/app /workspace/gopath/src/app
COPY --from=builder /workspace/gopath/src/app .
CMD ["./app"]