package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	// Host name of the HTTP Server
	Host = "localhost"
	// Port of the HTTP Server
	Port = "8000"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got / request")
	io.WriteString(w, "I am alive!\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}
func getWebhook(urlString *http.Request) {
	fmt.Println(" received /webhook request")
	fmt.Println(urlString)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.URL.Path != "/webhook" && r.URL.Path != "/hello" {
		http.Error(w, "404 not found. Only '/', '/hello' or '/webhook' are valid paths", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		http.HandleFunc("/", getRoot)
	case "POST":
		getWebhook(r)
	}
}

func startServer() {
	http.HandleFunc("/raba", handleRequest)
	http.HandleFunc("/hello", getHello)
	fmt.Printf("Starting server for testing on localhost:" + Port + "...\n")
	err := http.ListenAndServe(":"+Port, nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed\n")
	} else if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}
