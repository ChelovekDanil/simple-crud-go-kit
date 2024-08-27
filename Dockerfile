FROM golang:1.22.6 AS builder
WORKDIR /src
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o crud ./cmd/crud

FROM scratch
WORKDIR /app
COPY --from=builder /src/configs/config.yaml /app/config/
COPY --from=builder /src/init.sql ./init.sql:/docker-entrypoint-initdb.d/init.sql
COPY --from=builder /src/crud .
EXPOSE 8080
CMD ["./crud"]