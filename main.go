package main

import (
	"net/http"
	"os"
	"fmt"
	"log"
)

func handle(w http.ResponseWriter, r *http.Request) {
	log.Println("request recived, responding message...")
	fmt.Fprint(w, "hi, demo started")
}

func main() {
	port := os.Getenv("PORT")
	log.Println("Port:",port)
	http.HandleFunc("/", handle)
	log.Println("Starting service...")
	http.ListenAndServe(":"+port+"/", nil)
}
