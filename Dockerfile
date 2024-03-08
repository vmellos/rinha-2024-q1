FROM golang:1.22-alpine as builder
WORKDIR /
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o ./Api ./cmd/main.go

FROM alpine:latest
WORKDIR /
COPY --from=builder . .
ENTRYPOINT ["./Api"]