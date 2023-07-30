FROM golang:1.19-alpine as builder

WORKDIR /app

COPY ./ ./

RUN go mod download && go get -u ./...

RUN go mod vendor

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8081

CMD ["./main"]