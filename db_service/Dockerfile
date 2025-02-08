FROM golang:1.23.5
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o db_service main.go
CMD ["./db_service"]
