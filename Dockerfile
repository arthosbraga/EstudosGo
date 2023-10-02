FROM golang:1.16-alpine AS builder

WORKDIR /app

COPY ./src /app/
RUN go mod download
RUN go build -o ./server ./main.go

FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8000
ENTRYPOINT ["./server"]