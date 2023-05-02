ARG VARIANT=1.20-bullseye
FROM golang:${VARIANT}

# OS
RUN apt update

# Golang
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
