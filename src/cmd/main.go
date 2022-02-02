package main

import (
	"log"
	"net/http"

	"github.com/KenethSandoval/xopopu/internal/router"
)

func main() {
	mux := http.NewServeMux()

	hs := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	router.InitRouter(mux)

	if err := hs.ListenAndServe(); err != nil {
		log.Fatalf("%v", err)
	}
}
