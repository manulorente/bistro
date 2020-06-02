package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// The new router function creates the router and
// returns it to us. We can now use this function
// to instantiate and test the router outside of the main function
func CreateRouter() *mux.Router {
	r := mux.NewRouter()

	// Create a file server which serves files out of the "./ui/static" directory.
	// Note that the path given to the http.Dir function is relative to the project
	// directory root.
	fileServer := http.FileServer(http.Dir("../frontend"))
	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	r.Handle("/frontend/", http.StripPrefix("/frontend/", fileServer))

	// Declare the static file directory and point it to the
	// directory we just made
	staticFileDirectory := http.Dir("../frontend/views/")
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

	// Views
	r.HandleFunc("/", IndexHandler).Methods("GET")

	r.HandleFunc("/menu", GetMenuHandler).Methods("GET")
	r.HandleFunc("/menu", CreateMenuHandler).Methods("POST")

	r.HandleFunc("/products", GetMenuEndPoint).Methods("GET")
	r.HandleFunc("/products/{id}", GetProductEndPoint).Methods("GET")
	r.HandleFunc("/products/{id}", CreateProductEndPoint).Methods("POST")
	r.HandleFunc("/products/{id}", DeleteProductEndPoint).Methods("DELETE")
	r.HandleFunc("/products/{id}", UpdateProductEndPoint).Methods("PUT")

	// Just to test
	products = append(products, Product{ID: "1", Description: "Cerveza", Size: "Caña", Price: "1€"})
	products = append(products, Product{ID: "2", Description: "Montadito", Size: "2uds", Price: "3€"})

	return r
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../frontend/views/index.html")
	log.Print("Index page")
}
