package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ServerController struct {
	session *mgo.Session
}

func NewServerController(s *mgo.Session) *ServerController {
	return &ServerController{s}
}

func(sc ServerController) GetAllResponses(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	enableCors(&w)
	var data []bson.M
	if err := sc.session.DB("so1-test").C("so1-test").Find(bson.M{}).All(&data); err != nil {
		w.WriteHeader(404)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json, _ := json.Marshal(data)
	fmt.Fprintf(w, "%s\n", json)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
