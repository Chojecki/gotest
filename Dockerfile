FROM golang:1.14.9-alpine AS builder
RUN mkdir /build
ADD go.mod go.sum hello.go /build/
WORKDIR /build
RUN go build

FROM alpine
RUN adduser -S -D -H -h /app omgimjedi
USER omgimjedi
COPY --from=builder /build/gotest /app/
COPY views/ /app/views
WORKDIR /app
CMD ["./gotest"]