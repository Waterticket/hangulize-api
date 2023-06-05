FROM golang:1.20.4 AS builder

RUN mkdir -p /build
WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod vendor
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o bin/hangulize-api ./cmd/server.go

RUN mkdir -p /dist
WORKDIR /dist
RUN cp /build/bin/hangulize-api ./hangulize-api



FROM golang:1.20.4-alpine3.18

RUN mkdir -p /app
WORKDIR /app

COPY --chown=0:0 --from=builder /dist /app/
EXPOSE 5000

ENTRYPOINT ["/app/hangulize-api"]
