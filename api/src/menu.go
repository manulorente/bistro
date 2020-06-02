package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Menu : Estructura json para intercambio de mensajes
type Menu struct {
	Description string `json:"description,omitempty"`
	Size        string `json:"size,omitempty"`
	Price       string `json:"price,omitempty"`
}

// Menu : Estructura json para intercambio de mensajes
type Product struct {
	ID          string `json: "id,omitempty"`
	Description string `json:"description,omitempty"`
	Size        string `json:"size,omitempty"`
	Price       string `json:"price,omitempty"`
}

// Here will go database conn
var tapas []Menu
var products []Product

func GetMenuHandler(w http.ResponseWriter, r *http.Request) {
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
	//w.Header().Set("Content-Type", "application/json")
	w.Write(menuListBytes)
	log.Print("getMenu Handler")
}

func CreateMenuHandler(w http.ResponseWriter, r *http.Request) {
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
	tapa.Description = r.Form.Get("description")
	tapa.Size = r.Form.Get("size")
	tapa.Price = r.Form.Get("price")

	// Append our existing list of tapas with a new entry
	tapas = append(tapas, tapa)

	//Finally, we redirect the user to the original HTMl page
	// (located at `/views/`), using the http libraries `Redirect` method
	http.Redirect(w, r, "/views/", http.StatusFound)
	log.Print("createMenu Handler")
}

func GetMenuEndPoint(w http.ResponseWriter, r *http.Request) {
	// Encoder is used to pass objects
	json.NewEncoder(w).Encode(products)
	log.Print("getMenu End Point")
}

func GetProductEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range products {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			log.Print("getProduct End Point found")
			return
		}
	}
	// We return empty product if it is not found
	json.NewEncoder(w).Encode(&Product{})
	log.Print("getProduct End Point not found")
}

func CreateProductEndPoint(w http.ResponseWriter, r *http.Request) {
	// We need to read from the URL
	params := mux.Vars(r)
	var product Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	product.ID = params["id"]
	products = append(products, product)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(products)
	log.Print("CreateProduct End Point")
}

func DeleteProductEndPoint(w http.ResponseWriter, r *http.Request) {
	// We need to delete from the URL
	params := mux.Vars(r)
	for index, item := range products {
		if item.ID == params["id"] {
			products = append(products[:index], products[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(products)
	log.Print("DeleteProduct End Point")
}

func UpdateProductEndPoint(w http.ResponseWriter, r *http.Request) {
	// We need to modify from the URL
	params := mux.Vars(r)
	var updatedProduct Product
	json.NewDecoder(r.Body).Decode(&updatedProduct)
	for index, item := range products {
		if item.ID == params["id"] {
			item.Description = updatedProduct.Description
			item.Size = updatedProduct.Size
			item.Price = updatedProduct.Price
			products = append(products[:index], item)
			json.NewEncoder(w).Encode(products)
			log.Print("UpdateProduct End Point")
			break
		}
	}
}
