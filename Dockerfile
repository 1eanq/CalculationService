FROM golang:1.19

WORKDIR /app

COPY . .
RUN go mod download

RUN go build -o ./calculator cmd/main.go

EXPOSE 8080
CMD ["./calculator"]