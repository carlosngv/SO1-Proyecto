package controllers

import (
	"carlosngv/so1-proyecto/models"
	"encoding/json"
	"fmt"
	"net/http"

	"os/exec"

	"github.com/julienschmidt/httprouter"
)

func ExecLocust (w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	enableCors(&w)

	fmt.Println("Locust executed!")
	app := "sh"

    arg1 := "/Users/carlosngv/Documents/U/SO1/Proyecto/loadTester/GoServer/loadData.sh"

    cmd := exec.Command(app, arg1)
    stdout, err := cmd.CombinedOutput()

    if err != nil {
        fmt.Println("[ERROR] " + err.Error())
        return
    }

    fmt.Println(string(stdout))
}

func SendData (w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	enableCors(&w)

	newResponse := models.Response{}

	json.NewDecoder(r.Body).Decode(&newResponse)

	uj, err := json.Marshal(newResponse)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)

}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
