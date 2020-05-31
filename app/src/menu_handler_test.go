package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)

func TestGetMenuHandler(t *testing.T) {

	tapas = []Menu{
		{"Ensaladilla", "5€"},
	}

	req, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(getMenuHandler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := Menu{"Ensaladilla", "5€"}
	b := []Menu{}
	err = json.NewDecoder(recorder.Body).Decode(&b)

	if err != nil {
		t.Fatal(err)
	}

	actual := b[0]

	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}
func TestCreateMenuHandler(t *testing.T) {

	tapas = []Menu{
		{"Ensaladilla", "5€"},
	}

	form := newCreateMenuForm()
	req, err := http.NewRequest("POST", "", bytes.NewBufferString(form.Encode()))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(createMenuHandler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := Menu{"Montadito", "2€"}

	if err != nil {
		t.Fatal(err)
	}

	actual := tapas[1]

	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func newCreateMenuForm() *url.Values {
	form := url.Values{}
	form.Set("tapas", "Montadito")
	form.Set("precio", "2€")
	return &form
}
