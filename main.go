package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling root request")
	fmt.Fprintf(w, time.Now().String()) // send data to client side
	log.Println("Request completed")
}

func main() {
	http.HandleFunc("/", sayhelloName)       // set router
	err := http.ListenAndServe(":9091", nil) // set listen port
	log.Println("App listening on 9091")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
