package main

import (
	// System
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"time"

	// Local

	// Third party
	"github.com/gorilla/mux"
)

// Port: Puerto de entrada al servidor
const (
	Port = ":8080"
)

// Estructura que se usa para el objeto servidor
type message struct {
	msg string
}

// Product : La API REST devolverá contenido JSON que permitirá que el frontend sea independiente al backend
type Product struct {
	Cat         string    `json:"cat,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Size        string    `json:"size,omitempty"`
	Price       string    `json:"price,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

// Menu :  de un restaurante
type Menu struct {
	Language string
	Products []Product
}

// Restaurant :
type Restaurant struct {
	Name  string
	Menus []Menu
}

// En producción debemos añadir paginacion para devolver objetos de un solo usuario

// Creamos map como BBDD virtual
var productStore = make(map[string]Product)

// Puntero para escribir/leer de la BBDD virtual
var id int

// GetProductsHandler - GET - /products
func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	// Debemos recorrer el mapa virtual para almacenar en la BBDD virtual
	var products []Product
	for _, value := range productStore {
		products = append(products, value)
	}
	// Codificamos el dato para que sea una estructura de json
	j, err := json.Marshal(products)
	if err != nil {
		log.Print(err)
	}
	// Preparamos la trama [Tipo - Cabecera - Cuerpo]
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
	log.Print(getFnName() + " " + strconv.Itoa(http.StatusOK))
}

// PostProductsHandler - POST - /products
func PostProductsHandler(w http.ResponseWriter, r *http.Request) {
	// Decodificamos el dato del cuerpo de la peticion y lo mapeamos a la estructura
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Print(err)
	}
	// Agregamos fecha de creacion
	product.CreatedAt = time.Now()
	// Incrementamos indice de la BBD
	id++
	k := strconv.Itoa(id)
	// Almacenamos en la BBDD
	productStore[k] = product
	// Devolvemos objeto creado en formato JSON
	j, err := json.Marshal(product)
	if err != nil {
		log.Print(err)
	}
	// Preparamos la trama [Tipo - Cabecera - Cuerpo]
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
	log.Print(getFnName() + " " + strconv.Itoa(http.StatusCreated))
}

// PutProductsHandler - PUT - /products/id
func PutProductsHandler(w http.ResponseWriter, r *http.Request) {
	// Extraemos la variable y el id del objeto a actualizar
	vars := mux.Vars(r)
	k := vars["id"]
	// Decodificamos el dato del cuerpo de la peticion y lo mapeamos a la estructura
	var productUpdate Product
	err := json.NewDecoder(r.Body).Decode(&productUpdate)
	if err != nil {
		log.Print(err)
	}
	// Buscamos objeto y si existe hacemos el cambio (eliminar - insertar)
	if product, ok := productStore[k]; ok {
		productUpdate.CreatedAt = product.CreatedAt
		delete(productStore, k)
		productStore[k] = productUpdate
	} else {
		log.Printf("ID %s not found ", k)
	}

	w.WriteHeader(http.StatusNoContent)
	log.Print(getFnName() + " " + strconv.Itoa(http.StatusNoContent))
}

// DelProductsHandler - DELETE - /products/id
func DelProductsHandler(w http.ResponseWriter, r *http.Request) {
	// Extraemos la variable y el id del objeto a actualizar
	vars := mux.Vars(r)
	k := vars["id"]

	// Buscamos objeto y si existe lo eliminamos
	if _, ok := productStore[k]; ok {
		delete(productStore, k)
	} else {
		log.Printf("ID %s not found ", k)
	}

	w.WriteHeader(http.StatusNoContent)
	log.Print(getFnName() + " " + strconv.Itoa(http.StatusNoContent))
}

// Implementamos metodo de servidor
func (m message) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.msg)
	log.Print(getFnName() + " " + strconv.Itoa(http.StatusOK))
}

// Punto de entrada del programa
func main() {
	// Creamos handler
	msg := message{msg: "<h1>Index page</h1>"}

	// Creamos objeto de routing - StrictSlah permite que "/view" y "/view/" sea lo mismo
	r := mux.NewRouter().StrictSlash(false)

	// Vamos a hacer uso de la funcion Handle para asociar un handler a la ruta indicada
	r.Handle("/", msg)

	// Parseamos estructura de datos de tipo texto en un template
	menuEsp := Menu{
		Language: "ESP",
		Products: []Product{
			{"Comida", "Caracoles", "De Triana", "Tarrina", "5€", time.Now()},
			{"Bebida", "Vino", "Rioja", "Botella", "12€", time.Now()},
			{"Postre", "Ensalada", "De fruta", "Unidad", "5€", time.Now()}}}
	menuEng := Menu{
		Language: "ENG",
		Products: []Product{
			{"Food", "Nails", "From Triana", "Plate", "5€", time.Now()},
			{"Drink", "Wine", "Rioja", "Bottle", "12€", time.Now()},
			{"Desert", "Salad", "Fresh fruit", "Unit", "5€", time.Now()}}}

	restaurant := Restaurant{
		Name: "Restaurante de ejemplo",
		Menus: []Menu{
			menuEsp,
			menuEng}}

	t := template.New("Restaurant")
	t, err := t.ParseGlob("templates/*.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = t.ExecuteTemplate(os.Stdout, "test", restaurant)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Creamos CRUD (Create - Read - Update - Delete) para la entidad users
	r.HandleFunc("/products", GetProductsHandler).Methods("GET")
	r.HandleFunc("/products", PostProductsHandler).Methods("POST")
	r.HandleFunc("/products/{id}", PutProductsHandler).Methods("PUT")
	r.HandleFunc("/products/{id}", DelProductsHandler).Methods("DELETE")

	// Creamos estructura server para poder customizarlo
	server := &http.Server{
		Addr:           Port,             // Puerto de escucha
		Handler:        r,                // Nombre del handler
		ReadTimeout:    10 * time.Second, // Tiempo de respuesta
		WriteTimeout:   10 * time.Second, // Tiempo de respuesta
		MaxHeaderBytes: 1 << 20,          // 1MB maximo como header
	}

	// Arrancamos servidor
	log.Print("Listenning at port " + server.Addr)
	log.Fatal(server.ListenAndServe())
}

func getFnName() string {
	// Skip this function, and fetch the PC and file for its parent
	pc, _, _, _ := runtime.Caller(1)
	// Retrieve a Function object this functions parent
	functionObject := runtime.FuncForPC(pc)
	// Regex to extract just the function name (and not the module path)
	extractFnName := regexp.MustCompile(`^.*\.(.*)$`)
	return extractFnName.ReplaceAllString(functionObject.Name(), "$1")
}
