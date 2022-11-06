package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	pb "grpc-client/confproto"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

const (
	address = "localhost:5050" // ruta del servidor
)

func connectServer(w http.ResponseWriter, req *http.Request) {
	// AGREGAR CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	if req.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"msg\": \"ok\"}"))
		return
	}

	datos, _ := ioutil.ReadAll(req.Body)

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	// conn, err := grpc.Dial(os.Getenv("GRPC_SERVER_ADDRESS") + ":5050", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		json.NewEncoder(w).Encode("Error connecting to GRPC server.")
		log.Fatalf("Error: %v", err)
	}

	defer conn.Close()

	cl := pb.NewGetInfoClient(conn)

	id := string(datos)
	if len(os.Args) > 1 {
		id = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	ret, err := cl.ReturnInfo(ctx, &pb.RequestId{Id: id})
	if err != nil {
		json.NewEncoder(w).Encode("Error connecting to GRPC server.")
		log.Fatalf("Error: %v", err)
	}

	log.Printf("Server response: %s\n", ret.GetInfo())
	json.NewEncoder(w).Encode("Data successfully saved.")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", connectServer)
	fmt.Println("Server running on port 8200")
	log.Fatal(http.ListenAndServe(":8200", router))
}
