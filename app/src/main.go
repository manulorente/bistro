// This is the name of our package
// Everything with this package name can see everything
// else inside the same package, regardless of the file they are in
package main

// These are the libraries we are going to use
// Both "fmt" and "net" are part of the Go standard library
import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// The new router function creates the router and
// returns it to us. We can now use this function
// to instantiate and test the router outside of the main function
func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")
	r.HandleFunc("/", handler).Methods("GET")

	// Declare the static file directory and point it to the
	// directory we just made
	staticFileDirectory := http.Dir("./views/")
	// Declare the handler, that routes requests to their respective filename.
	// The fileserver is wrapped in the `stripPrefix` method, because we want to
	// remove the "/views/" prefix when looking for files.
	// For example, if we type "/views/index.html" in our browser, the file server
	// will look for only "index.html" inside the directory declared above.
	// If we did not strip the prefix, the file server would look for
	// "./views/views/index.html", and yield an error
	staticFileHandler := http.StripPrefix("/views/", http.FileServer(staticFileDirectory))
	// The "PathPrefix" method acts as a matcher, and matches all routes starting
	// with "/views/", instead of the absolute route itself
	r.PathPrefix("/views/").Handler(staticFileHandler).Methods("GET")
	r.HandleFunc("/menu", getMenuHandler).Methods("GET")
	r.HandleFunc("/menu", createMenuHandler).Methods("POST")

	return r
}

func main() {
	// The router is now formed by calling the `newRouter` constructor function
	// that we defined above. The rest of the code stays the same
	r := newRouter()
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
