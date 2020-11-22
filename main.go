package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func getPort() string {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	return port

}

func APIHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello Goper")

}

func main() {

	port := getPort()
	http.HandleFunc("/", APIHandler)
	//  Start HTTP
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal("Failed starting http server: ", err)
	}

}
