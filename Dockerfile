FROM golang:1.22-alpine AS builder

WORKDIR /usr/src/app

#cache
COPY go.mod go.sum ./
RUN go mod download && go mod tidy

#build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o ./bin/app cmd/main.go

#image
FROM alpine
COPY --from=builder /usr/src/app/bin/app /

COPY resources/config/*.yaml  resources/config/
COPY .env /.env

CMD ["/app"]