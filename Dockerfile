# Build the Go binary
FROM golang:1.20 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
#RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o snippetbox .

# Copy the binary to a new image
FROM alpine:latest

WORKDIR /

COPY --from=build-stage /app/snippetbox snippetbox

CMD ["/snippetbox"]
