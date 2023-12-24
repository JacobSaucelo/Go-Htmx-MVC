package routes

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"webGoCrud/controllers"
)

var portNumber uint16 = 3000

func RoutesIndex() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", controllers.DisplayUsers)
	mux.HandleFunc("/add-user/", controllers.AddUser)
	mux.HandleFunc("/delete-user/", controllers.DeleteUser)

	port := strconv.Itoa(int(portNumber))

	fmt.Println("Server: Running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
