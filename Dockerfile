FROM golang:1.15

COPY . /src
WORKDIR /src
RUN go build -o /brother-exporter /src/main.go

ENTRYPOINT /brother-exporter