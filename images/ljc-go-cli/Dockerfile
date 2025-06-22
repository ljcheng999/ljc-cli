FROM --platform=${TARGETOS}/${TARGETARCH} golang:1.24.4-alpine AS builder

LABEL author="Lucious Johnny Cheng"

ENV CGO_ENABLED=0 \
    GO111MODULE=on \
    GOOS=linux

COPY go.mod go.sum ./
RUN go mod download


COPY . .

RUN go build -ldflags="-w -s" -o /usr/local/bin/ljc-deploy

CMD ["ljc-deploy"]