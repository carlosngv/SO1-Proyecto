package main

import (
	"carlosngv/so1-proyecto/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func main() {
	// app := "sh"
	// arg1 := "/Users/carlosngv/Documents/U/SO1/Proyecto/loadTester/GoServer/loadData.sh"

    // cmd := exec.Command(app, arg1)
    // stdout, err := cmd.CombinedOutput()

    // if err != nil {
    //     fmt.Println("[ERROR] " + err.Error())
    //     return
    // }

    //fmt.Println(string(stdout))
	r := httprouter.New()
	r.GET("/locust_data", controllers.ExecLocust)
	r.POST("/locust_data", controllers.SendData)
	handler := cors.Default().Handler(r)
	http.ListenAndServe(":9000", handler)
}
