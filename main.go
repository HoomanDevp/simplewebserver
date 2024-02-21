package main

import (
	"fmt"
	"log"
	"net/http"
)

const portName string = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homepage")
}
func Info(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Info")
}

func main() {
	log.Println("Started on port", portName)
	fmt.Println("To close connection CTRL+C =)")

	http.HandleFunc("/", Home)
	http.HandleFunc("/info", Info)

	err := http.ListenAndServe(portName, nil)
	if err != nil {
		log.Fatal(err)
	}
	//Hi
}
