package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, time.Now().String()) // send data to client side
}

func main() {
	http.HandleFunc("/", sayhelloName)       // set router
	err := http.ListenAndServe(":9091", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
