FROM golang:1.20-alpine as builder

WORKDIR /app
COPY . .

RUN apk update && apk add gcc libc-dev && go mod tidy && \ 
  go build -a -ldflags "-linkmode external -extldflags '-static' -s -w" -o  bin/domain-seller cmd/domain-seller/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/bin/domain-seller /app/.env ./
ADD ./web /app/web

ENV GIN_MODE=release

EXPOSE 8080

CMD ["./domain-seller"]

