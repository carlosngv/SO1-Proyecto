package main

import (
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"gopkg.in/mgo.v2"

	"go-mongo/controllers"
)

func main() {

	r := httprouter.New()

	sc := controllers.NewServerController(getSession())

	r.GET("/results", sc.GetAllResponses)
	handler := cors.Default().Handler(r)
	http.ListenAndServe(":9001", handler)

}

func getSession() *mgo.Session {
	// Mongo DB connection
	s, err := mgo.Dial(os.Getenv("MONGO_ADDRES"))
	// s, err := mgo.Dial(os.Getenv("mongodb://so1-mongodb:D6KMglupcEagiKMis6Fff213MAD63Yz3VNAxghsdIyQqADe6ch21EMfovPV9a2DlT0cXKzjb2gjvACDbbmd6HA==@so1-mongodb.mongo.cosmos.azure.com:10255/?ssl=true&replicaSet=globaldb&retrywrites=false&maxIdleTimeMS=120000&appName=@so1-mongodb@"))
	if err != nil {
		panic(err)
	}
	return s
}
