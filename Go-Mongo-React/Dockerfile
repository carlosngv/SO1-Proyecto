FROM golang:1.17-alpine

ENV MONGO_ADDRESS mongodb://so1-mongodb:D6KMglupcEagiKMis6Fff213MAD63Yz3VNAxghsdIyQqADe6ch21EMfovPV9a2DlT0cXKzjb2gjvACDbbmd6HA==@so1-mongodb.mongo.cosmos.azure.com:10255/?ssl=true&replicaSet=globaldb&retrywrites=false&maxIdleTimeMS=120000&appName=@so1-mongodb@

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 9001

CMD ["/app/main"]
