package main

import (
	"fmt"
	"log"
	"net/http"
	"syscall"
)

func main() {
	log.Println("Service is started")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello!")
	})
	if err := http.ListenAndServe("0.0.0.0:5000", nil); err != nil {
		log.Fatal(err)
	}

}
