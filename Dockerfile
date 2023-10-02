FROM golang:latest AS Build

WORKDIR /app

COPY ./src .
RUN go mod download
RUN go build -o ./build/main main.go
RUN ls

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./build/main
FROM scratch


COPY --from=Build ./app/build/main main
EXPOSE 8000
ENTRYPOINT [ "/main" ]