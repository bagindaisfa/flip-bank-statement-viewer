# build stage
FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bankviewer ./cmd/server

# final
FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=builder /bankviewer /bankviewer
EXPOSE 8080
ENTRYPOINT ["/bankviewer"]
