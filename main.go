package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

var httpListenAndServe = http.ListenAndServe

func main() {
	log.Printf("Starting the application\n")
	RunServer()
}

func RunServer() {
	log.Printf("Running the server\n")
	mux := http.NewServeMux()
	mux.HandleFunc("/", VersionServer)
	log.Fatal("ListenAndServe: ", httpListenAndServe(":8080", mux))
}

func VersionServer(w http.ResponseWriter, req *http.Request) {
	msg := fmt.Sprintf("%s request to %s\n", req.Method, req.RequestURI)
	io.WriteString(w, msg)
}
