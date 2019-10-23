FROM golang:1.13.3@sha256:6a693fbaba7dd8d816f6afce049fb92b280c588e0a677c4c8db26645e613fc15 as builder
WORKDIR /builder/
COPY . .

RUN CGO_ENABLED=0 \
  GOOS=linux \
  go build \
  -a \
  -installsuffix cgo \
  -o cfwctl \
  cmd/cfwctl/main.go

FROM alpine:3.10.3@sha256:c19173c5ada610a5989151111163d28a67368362762534d8a8121ce95cf2bd5a

RUN apk add --no-cache ca-certificates

WORKDIR /bin/
# Copy the Pre-built binary file from the previous stage
COPY --from=builder /builder/cfwctl .
EXPOSE 8080
CMD ["./cfwctl"]
