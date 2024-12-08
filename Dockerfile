FROM golang:1.22.6 AS builder
WORKDIR /src
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o crud ./cmd/crud

FROM scratch
WORKDIR /app
COPY --from=builder /src/crud .
EXPOSE 8080
CMD ["./crud"]