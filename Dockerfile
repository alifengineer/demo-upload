FROM golang:1.18-alpine

RUN mkdir api 

WORKDIR /api

COPY ./ ./

RUN go mod download
RUN go mod vendor
RUN go build -o main ./cmd/main.go

EXPOSE 8081

CMD ["./main"]