package main

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"encoding/json"
	"log"
	"os"
)

var images Images

type Images struct {
	Images []Image `json:"images"`
}

type Image struct {
	Id string `json:"id"`
	Text string `json:"text"`
	Description string `json:"description"`
	ImageURL string `json:"image_url"`
}

func imagesHandle(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(images, "", " ")
	if err != nil {
		log.Panic(err)
	}
	fmt.Fprint(w, string(bytes))
}

func init() {
	dataFile, err := ioutil.ReadFile("./data.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(dataFile, &images)
}

func main() {
	port := os.Getenv("PORT")
	router := mux.NewRouter()
	router.HandleFunc("/images", imagesHandle)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	http.Handle("/", router)
	http.ListenAndServe(":"+port+"/", nil)
}
