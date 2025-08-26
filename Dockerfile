# Stage 1: Build
FROM golang:1.23 AS builder

WORKDIR /app

# Copy go.mod dan go.sum
COPY ../go.mod ../go.sum ./
RUN go mod download

# Copy semua source code
COPY .. .

# Build binary statis
RUN cd project && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -buildvcs=false -o /app/main .

# Stage 2: Run
# pakai distroless base image, ringan & modern
FROM gcr.io/distroless/base-debian12

WORKDIR /root/

# Copy binary dan .env
COPY --from=builder /app/main .
COPY --from=builder /app/project/.env .env

EXPOSE 8080
CMD ["./main"]
