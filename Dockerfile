FROM golang:1.25-alpine AS builder

WORKDIR /app

# Add go.sum if have deps
COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/web/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/server .

EXPOSE 7777

CMD ["./server"]
