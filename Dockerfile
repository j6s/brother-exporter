FROM golang:1.15 as builder

COPY . /build
WORKDIR /build
RUN go build -o /brother-exporter /build/main.go


FROM alpine:3.13

RUN apk add --no-cache libc6-compat
COPY --from=builder /brother-exporter /usr/local/bin/brother-exporter

ENTRYPOINT brother-exporter
