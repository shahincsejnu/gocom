FROM golang:1.18-buster as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o server .

FROM debian:buster-slim as runner

RUN apt update && apt install -y ca-certificates

WORKDIR /app

COPY --from=builder app/server .

EXPOSE 8080

CMD ["./server"]