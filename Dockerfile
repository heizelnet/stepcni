# Build the manager binary
FROM golang:alpine as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY ./ ./

# Build
# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -mod=vendor -o bin/mycnid cmd/mycnid/main.go && \
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -mod=vendor -o bin/mycni controller/main.go

FROM alpine
RUN apk update && apk add --no-cache iptables
WORKDIR /
COPY --from=builder /workspace/bin/* /

