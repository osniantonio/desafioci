FROM golang:1.12-alpine as builder
WORKDIR /src/app
ADD src/app /src/app
COPY . .
# para a imagem reduzida
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w"

FROM scratch
WORKDIR /src/app
ADD src/app /src/app
COPY --from=builder /src/app .
CMD ["./app"]