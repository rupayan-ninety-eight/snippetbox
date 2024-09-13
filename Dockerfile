# Build the Go binary
FROM golang:1.23.1 AS build-stage

WORKDIR /build
COPY go.mod go.sum ./
#RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o ./snippetbox

# Copy the binary to a new image
FROM alpine:latest
WORKDIR /app
COPY --from=build-stage /build/snippetbox ./snippetbox

# Expose the port the app runs on
EXPOSE 4000

CMD ["/app/snippetbox"]
