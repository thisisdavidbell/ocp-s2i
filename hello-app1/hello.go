package main

import (
	"fmt"
	"net/http"
)

// PORT - server port
var PORT = "8080"

// PATH - server path
var PATH = "/hello"

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello World, I'm app1. Boo!.\n")
}

func main() {
	fmt.Printf("Running hello server on %s:%s\n", PATH, PORT)
	http.HandleFunc(PATH, hello)
	http.ListenAndServe(":"+PORT, nil)
}
