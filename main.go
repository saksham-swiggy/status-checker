package main

import (
	"github.com/go-chi/chi"
	"github.com/saksham-swiggy/status-checker/handler"
	"log"
	"net/http"
)

func main() {
	router := chi.NewRouter()
	router.Post("/websites", handler.AddWebsites)
	router.Get("/websites", handler.GetStatus)
	log.Print("Starting HTTP server....")
	err := http.ListenAndServe("localhost:3000", router)
	if err != nil {
		log.Print(err)
	}
}
