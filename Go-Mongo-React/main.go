package main

import (
	"context"
	"os"

	"net/http"

	"encoding/json"

	_ "log"

	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"

	"go.mongodb.org/mongo-driver/mongo/options"
)



func main() {

     http.HandleFunc("/results", requestHandler)

     http.ListenAndServe(":9001", nil)

}



func requestHandler(w http.ResponseWriter, req *http.Request) {



    w.Header().Set("Content-Type", "application/json")



    response := map[string]interface{}{}



    ctx := context.Background()



    // client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://so1-mongodb:D6KMglupcEagiKMis6Fff213MAD63Yz3VNAxghsdIyQqADe6ch21EMfovPV9a2DlT0cXKzjb2gjvACDbbmd6HA==@so1-mongodb.mongo.cosmos.azure.com:10255/?ssl=true&replicaSet=globaldb&retrywrites=false&maxIdleTimeMS=120000&appName=@so1-mongodb@"))
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_ADDRESS")))



    if err != nil {

        fmt.Println(err.Error())

    }



    collection := client.Database("so1-test").Collection("so1-test")



    data := map[string]interface{}{}



    err = json.NewDecoder(req.Body).Decode(&data)



    if err != nil {

        fmt.Println(err.Error())

    }



    switch req.Method {

        case "GET":

            response, err = getRecords(collection, ctx)


    }



    if err != nil {

        response = map[string]interface{}{"error": err.Error(),}

    }

    enc := json.NewEncoder(w)
    enc.SetIndent("", "  ")

    if err := enc.Encode(response); err != nil {
        fmt.Println(err.Error())
    }
}

func getRecords(collection *mongo.Collection, ctx context.Context)(map[string]interface{}, error){

    cur, err := collection.Find(ctx, bson.M{})

    if err != nil {
        return nil, err
    }

    defer cur.Close(ctx)
    var results []bson.M

    for cur.Next(ctx) {
        var product bson.M
        if err = cur.Decode(&product); err != nil {
            return nil, err
        }

        results = append(results, product)
    }

    res := map[string]interface{}{}
    res = map[string]interface{}{
              "data" : results,
          }



    return res, nil

}
