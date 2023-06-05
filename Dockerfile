FROM golang:1.20.4 AS builder

RUN mkdir -p /build/bin
WORKDIR /build

COPY go.mod .
COPY go.sum .
COPY . .
RUN go build -o bin/hangulize-api ./main.go

RUN mkdir -p /dist
WORKDIR /dist
RUN cp /build/bin/hangulize-api ./hangulize-api



FROM ubuntu:20.04

RUN mkdir -p /app
WORKDIR /app

COPY --chown=0:0 --from=builder /dist /app
RUN chmod a+x /app/hangulize-api
RUN ls -al
EXPOSE 5000

ENTRYPOINT ["/app/hangulize-api"]
