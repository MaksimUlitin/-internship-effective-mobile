
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o music-library-app ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/music-library-app .
COPY --from=builder /app/enrichInfoSong.json .
COPY --from=builder /app/.env .

EXPOSE 8080
EXPOSE 8088

CMD ["./music-library-app"]