FROM golang:1.21-alpine3.18 AS builder

RUN mkdir /app
ADD . /app
COPY .env /app/
WORKDIR /app
RUN go clean --modcache
RUN go build -o main

FROM alpine:3.18
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/.env .
EXPOSE 8080
CMD ["./main"]