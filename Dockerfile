FROM golang:1.20-alpine as builder

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o bin/domain-seller cmd/domain-seller/main.go


FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/bin/domain-seller /app/.env ./
RUN mkdir -p /app/web/templates
COPY ./web/templates/* /app/web/templates
ENV GIN_MODE=release

EXPOSE 8080

CMD ["./domain-seller"]

