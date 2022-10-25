package main

import (
	"net/http"

	"carlosngv/so1-proyecto/controllers"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func main() {
	r := httprouter.New()
	r.GET("/locust_data", controllers.ExecLocust)
	r.POST("/locust_data", controllers.SendData)
	handler := cors.Default().Handler(r)
	http.ListenAndServe(":9000", handler)
}
