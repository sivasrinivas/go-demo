package main

import (
	"net/http"
	"fmt"
	"log"
	"github.com/gorilla/mux"
	"os"
)

func handle(w http.ResponseWriter, r *http.Request) {
	log.Println("request recived, responding message...")
	fmt.Fprint(w, "hi, demo started")
}

func main() {
	port := os.Getenv("PORT")
	router := mux.NewRouter()
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	http.Handle("/", router)
	http.ListenAndServe(":"+port+"/", nil)
}
