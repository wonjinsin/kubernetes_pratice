FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o main main.go

CMD ["./main"]