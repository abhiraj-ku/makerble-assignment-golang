FROM golang:1.22-alpine AS builder

RUN mkdir /app

COPY go.mod go.sum ./

RUN go mod download 

WORKDIR /app

COPY . .

RUN CGO_ENABLED= false go build -o heathcare-service ./cmd/health-app


# stage 2 

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app .

EXPOSE 9090

CMD ["./healthcare-service"]