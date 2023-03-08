FROM golang:1.20-alpine AS builder
WORKDIR /build
# download dependencies
COPY go.mod go.sum ./
RUN go mod download
# build binary
COPY . ./
RUN go build -o silly-gorest .

# Create final image
FROM alpine:latest
WORKDIR /
COPY --from=builder /build/silly-gorest .
EXPOSE 8000
ENTRYPOINT [ "./silly-gorest" ]