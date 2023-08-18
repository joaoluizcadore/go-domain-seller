FROM golang:1.20-alpine as builder

WORKDIR /app
COPY . .

RUN apk update && apk add gcc libc-dev
RUN go mod tidy && \ 
  go build -a -ldflags "-linkmode external -extldflags '-static' -s -w" -o  bin/domain-seller cmd/domain-seller/main.go


FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/bin/domain-seller /app/.env ./
RUN mkdir -p /app/web/templates && mkdir -p /app/web/static
COPY ./web/templates/* /app/web/templates
COPY ./web/static/ /app/web/static


ENV GIN_MODE=release

EXPOSE 8080

CMD ["./domain-seller"]

