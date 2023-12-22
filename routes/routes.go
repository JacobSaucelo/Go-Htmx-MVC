package routes

import (
	"log"
	"net/http"
	"webGoCrud/controllers"
)

func RoutesIndex() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", controllers.DisplayUsers)

	log.Fatal(http.ListenAndServe(":3000", mux))
}
