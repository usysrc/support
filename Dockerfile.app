FROM golang:1.20-alpine as builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o support cmd/support/main.go

FROM scratch

EXPOSE 8080

WORKDIR /

COPY --from=builder /app /usr/bin/app

CMD ["/usr/bin/app/support"]