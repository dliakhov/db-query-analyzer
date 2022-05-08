# build app
FROM golang:1.18-alpine3.15 as builder

RUN mkdir /build
WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o db-query-analyzer .

FROM alpine:3.15

COPY --from=builder /build/db-query-analyzer /app/db-query-analyzer
RUN mkdir /migrations
COPY ./migrations /migrations

ENTRYPOINT ["/app/db-query-analyzer", "httpservice"]
