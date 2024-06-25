# Build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git

WORKDIR /go/src/app
COPY . .

# Download dependencies
RUN go get -d -v ./...

# Build the application targeting the main package
RUN go build -o /go/bin/app .

# Final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates

COPY --from=builder /go/bin/app /app

ENTRYPOINT ["/app"]

LABEL Name=backendproject Version=0.0.1
EXPOSE 8080
