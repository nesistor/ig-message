# Use Golang image to build the application
FROM golang:1.20 AS builder
WORKDIR /src
COPY . .
RUN go mod tidy
RUN go build -o /app/main .

# Start a new image to serve the app
FROM debian:bullseye-slim
COPY --from=builder /app/main /app/main
WORKDIR /app
EXPOSE 8080
CMD ["./main"]
