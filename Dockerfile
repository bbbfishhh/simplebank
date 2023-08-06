# Build stage
FROM golang:1.20.6-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
# RUN apk add --no-cache curl
# RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar -xvz

# Run stage
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
# COPY --from=builder /app/migrate ./migrate
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
RUN chmod +x /app/start.sh
RUN chmod +x /app/wait-for.sh
COPY db/migration ./db/migration

EXPOSE 8888
EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]