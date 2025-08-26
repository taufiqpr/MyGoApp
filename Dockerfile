# Stage 1: Build
FROM golang:1.23 AS builder

WORKDIR /app

COPY ../go.mod ../go.sum ./
RUN go mod download

COPY .. .

RUN cd project && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -buildvcs=false -o /app/main .

FROM gcr.io/distroless/base-debian12

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/project/.env .env

EXPOSE 8080
CMD ["./main"]
