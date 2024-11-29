# Stage 1: Build
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY main.go .
RUN go build -o server main.go

# Stage 2: Run
FROM scratch
ENV TEXT="Hello, World!"
ENV PORT=8080
COPY --from=builder /app/server /server
EXPOSE 8080
ENTRYPOINT ["/server"]

