package router

import (
	"encoding/json"
	"net/http"
)

type Handler struct{}

func (Handler) ServerHTTP(w http.ResponseWriter, r *http.Request) {

}

func InitRouter(mux *http.ServeMux) {
	mux.HandleFunc("/", home)
}

func home(w http.ResponseWriter, r *http.Request) {
	type MessageStruct struct {
		Message string `json:"message"`
	}
	json.NewEncoder(w).Encode(MessageStruct{Message: "hello"})
}
