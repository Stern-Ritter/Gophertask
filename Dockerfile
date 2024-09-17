FROM golang:1.22 AS builder

WORKDIR /app/gophertask

COPY go.mod ./
RUN go mod tidy
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /gophertask ./cmd/server/main.go

FROM scratch

WORKDIR /app

COPY --from=builder /gophertask /app/gophertask
COPY --from=builder /app/gophertask/certs /app/certs

CMD ["/app/gophertask"]