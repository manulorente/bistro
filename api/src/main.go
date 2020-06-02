package main

import (
	"fmt"
	"log"
	"net/http"
)

// Page: Registros que guardan páginas
type message struct {
	msg string
}

func main() {
	// The router is now formed by calling the `newRouter` constructor function
	// that we defined above. The rest of the code stays the same
	r := CreateRouter()
	log.Print("Listenning at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func (m message) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, m.msg)
}
