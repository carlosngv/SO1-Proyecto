FROM golang:latest

WORKDIR /
COPY . .

RUN go mod download

EXPOSE 5050

CMD ["go", "run", "main.go"]
