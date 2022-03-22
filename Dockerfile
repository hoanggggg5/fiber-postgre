FROM golang:1.16-alpine3.15

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build main.go

EXPOSE 3000

CMD ["./main"]  