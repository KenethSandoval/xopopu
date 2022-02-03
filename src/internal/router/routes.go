package router

import (
	"encoding/json"
	"net/http"

	"github.com/KenethSandoval/xopopu/internal/upload"
)

type Handler struct{}

type MessageStruct struct {
	Message string `json:"message"`
}

func (Handler) ServerHTTP(w http.ResponseWriter, r *http.Request) {
}

func InitRouter(mux *http.ServeMux) {
	mux.HandleFunc("/", home)
	mux.HandleFunc("/upload", upload.UploadCSV)
}

func home(w http.ResponseWriter, _ *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(MessageStruct{Message: "hello"})
}
