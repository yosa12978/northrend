FROM golang:1.20-alpine3.17 as builder

WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o northrend ./cmd/northrend/main.go

FROM alpine:3.17

WORKDIR /app
COPY --from=builder /app/northrend .
COPY --from=builder /app/config.json .
COPY --from=builder /app/templates templates
COPY --from=builder /app/static static

EXPOSE 5000
ENTRYPOINT ["/app/northrend"]
