package handler

import (
	"encoding/json"
	"fmt"
	"hello/db"
	"hello/error"
	"hello/model"
	"io"
	"io/ioutil"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func Quote(w http.ResponseWriter, r *http.Request) {
	q := db.GetQuote()

	if q != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(q); err != nil {
			panic(err)
		}
		return
	}

	// 404 if not found
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(error.JsonError{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}

/*
Test
curl -H "Content-Type: application/json" -d '{"quote": "text", "category": "cat4"}' http://localhost:8080/new
*/
func New(w http.ResponseWriter, r *http.Request) {

	//fmt.Println(id)

	var quote model.Quote
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &quote); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	id := db.InsertQuote(&quote)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(id); err != nil {
		panic(err)
	}
}
