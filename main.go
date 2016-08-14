package main

import (
	"net/http"
	"os"
	"fmt"
	"log"
)

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hi, demo started")
}

func main() {
	port := os.Getenv("PORT")
	log.Println("port:",port)
	http.HandleFunc("/", handle)
	http.ListenAndServe(":"+port+"/", nil)
}
