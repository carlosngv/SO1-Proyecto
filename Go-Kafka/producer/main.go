package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/segmentio/kafka-go"
)

type Data struct {
	Team1   string `json:"team1"`
	Team2   string `json:"team2"`
	Score   string `json:"score"`
	Phase   string `json:"phase"`
}

func write(w http.ResponseWriter, request *http.Request) {
	w.Header().Add("content-type", "application/json")
	var data Data
	json.NewDecoder((request.Body)).Decode(&data)

	fmt.Printf("Data: %v\n\n" , data)
	// fmt.Printf("Host: %v\n\n" , os.Getenv("HOSTADDR"))

	// conn, _ := kafka.DialLeader(context.Background(), "tcp", os.Getenv("HOSTADDR") + ":9092", "proyecto", 0)
	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "proyecto", 0)
	conn.SetWriteDeadline(time.Now().Add(time.Second * 10))
	conn.WriteMessages(kafka.Message{Value: []byte("{\"team1\":\"" + data.Team1 + "\",\"team2\": \"" + data.Team2 + "\",\"Score\": \"" + data.Score + "\",\"phase\": \"" + data.Phase + "\"}")})

	json.NewEncoder(w).Encode("Escribiendo en kafka")

	resData, err := json.Marshal(data)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", resData)

}

func salute(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", "hello")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/input", write).Methods("POST")
	router.HandleFunc("/", salute).Methods("GET")
	fmt.Println("Server on port", 8000)
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		fmt.Println(err)
	}
}
