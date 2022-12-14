package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gorilla/mux"
	"github.com/segmentio/kafka-go"
)

func saveResponse(response string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://so1-mongodb:D6KMglupcEagiKMis6Fff213MAD63Yz3VNAxghsdIyQqADe6ch21EMfovPV9a2DlT0cXKzjb2gjvACDbbmd6HA==@so1-mongodb.mongo.cosmos.azure.com:10255/?ssl=true&replicaSet=globaldb&retrywrites=false&maxIdleTimeMS=120000&appName=@so1-mongodb@"))
	// mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_ADDRESS")))
	if err != nil {
		log.Fatal(err)
	}

	databases, err := mongoClient.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(databases)

	Database := mongoClient.Database("so1-test")
	Collection := Database.Collection("so1-test")

	var bdoc interface{}

	errb := bson.UnmarshalExtJSON([]byte(response), true, &bdoc)
	fmt.Println(errb)

	insertResult, err := Collection.InsertOne(ctx, bdoc)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data successfully saved: \n")
	fmt.Println(insertResult)
}

func consumeData(w http.ResponseWriter, request *http.Request) {

	fmt.Println("Consuming data")

	// conn, _ := kafka.DialLeader(context.Background(), "tcp", os.Getenv("HOSTADDR") + ":9092", "proyecto", 0)
	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "proyecto", 0)
	conn.SetWriteDeadline(time.Now().Add(time.Second * 3))

	batch := conn.ReadBatch(1e3, 1e6) // 1 - 1000
	bytes := make([]byte, 1e3)
	for {
		_, err := batch.Read(bytes)
		if err != nil {
			break
		}
		fmt.Println(string(bytes))
		saveResponse(string(bytes))

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", consumeData).Methods("GET")
	fmt.Println("Server on port", 8100)
	err := http.ListenAndServe(":8100", router)
	if err != nil {
		fmt.Println(err)
	}

}
