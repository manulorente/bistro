package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Menu struct {
	Tapa   string `json:"tapa"`
	Precio string `json:"precio"`
}

var tapas []Menu

func getMenuHandler(w http.ResponseWriter, r *http.Request) {
	//Convert the "tapas" variable to json
	menuListBytes, err := json.Marshal(tapas)

	// If there is an error, print it to the console, and return a server
	// error response to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// If all goes well, write the JSON list of tapas to the response
	w.Write(menuListBytes)
}

func createMenuHandler(w http.ResponseWriter, r *http.Request) {
	// Create a new instance of Menu
	tapa := Menu{}

	// We send all our data as HTML form data
	// the `ParseForm` method of the request, parses the
	// form values
	err := r.ParseForm()

	// In case of any error, we respond with an error to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get the information about the tapa from the form info
	tapa.Tapa = r.Form.Get("tapa")
	tapa.Precio = r.Form.Get("precio")

	// Append our existing list of tapas with a new entry
	tapas = append(tapas, tapa)

	//Finally, we redirect the user to the original HTMl page
	// (located at `/views/`), using the http libraries `Redirect` method
	http.Redirect(w, r, "/views/", http.StatusFound)
}
