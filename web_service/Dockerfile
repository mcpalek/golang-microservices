FROM golang:1.23.5
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o web_service main.go
EXPOSE 8080
CMD ["./web_service"]
