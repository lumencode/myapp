package main

import (
	"fmt"
	"log"
	"net/http"
)

func main () {

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("AWS"))
	})

	fmt.Println("Server readt at :4000")

	if err := http.ListenAndServe(":4000", nil); err != nil {
		log.Fatal(err.Error())
	}
}
