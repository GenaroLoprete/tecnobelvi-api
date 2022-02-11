package server

import (
	"log"
	"net/http"
)

func Init() {
	handleRequests()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRequests() {
	http.HandleFunc("/ping", Ping)
	http.HandleFunc("/builds", GetBuilds)
}
