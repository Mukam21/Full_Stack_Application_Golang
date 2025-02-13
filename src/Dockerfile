FROM golang:1.23.6-alpine3.20 AS builder

WORKDIR /go/src/app

ENV GO111MODULE=on

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./run .

FROM alpine:3.20
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /go/src/app/run .

EXPOSE 8080
CMD ["./run"]