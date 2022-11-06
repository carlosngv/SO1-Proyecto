package main

import (
	"context"
	"fmt"
	pb "grpc-server/confproto"
	"log"
	"net"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

const (
	port = ":5050"
)

type Response struct {
	Team1	string			`json:"team1"`
	Team2	string			`json:"team2"`
	Score	string			`json:"score"`
	Phase	string			`json:"phase"`
}

type Server struct {
	pb.UnimplementedGetInfoServer
}

func save_response(response string) {
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
	fmt.Println(insertResult)
}

func (s *Server) ReturnInfo(ctx context.Context, in *pb.RequestId) (*pb.InfoReply, error) {
	save_response(in.GetId())
	fmt.Printf(">> Response delivered: %v\n", in.GetId())
	return &pb.InfoReply{Info: ">> Response received by client: " + in.GetId()}, nil
}

func main() {

	listener, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterGetInfoServer(s, &Server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
