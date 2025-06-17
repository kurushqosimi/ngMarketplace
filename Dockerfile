FROM golang:1.24-alpine as BUILDER

WORKDIR /app

RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/app/main.go

FROM alpine

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/migrate.linux-amd64 /app/migrate
COPY /config/config.yaml /app/config/config.yaml

COPY wait-for.sh /app/wait-for.sh
COPY start.sh /app/start.sh

RUN chmod +x /app/start.sh /app/wait-for.sh

COPY migrations /app/migrations

RUN ls -la /app && ls -la /app/migrations
RUN apk add netcat-openbsd

EXPOSE 8081
CMD ["/app/main"]
ENTRYPOINT ["/app/start.sh"]
