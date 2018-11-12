package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ClientIdentity struct {
	id string
}

// Identity - is identity struct
type Identity struct {
	Key  string
	Type string
}

func main() {
	http.HandleFunc("/transaction", transactionHandler)
	http.ListenAndServe(":8080", nil)
}

func transactionHandler(w http.ResponseWriter, r *http.Request) {

	authentification()

	authorization()

	if r.Method != "POST" {
		http.Error(w, "Invalid request method", 544)
		return
	}

	var idn Identity
	err := json.NewDecoder(r.Body).Decode(&idn)
	fmt.Println(err)
	fmt.Println(idn)
	idn.Type = r.Method
	json.NewEncoder(w).Encode(idn)
}

func authentification() ClientIdentity {
	var ci ClientIdentity
	return ci
}

func authorization() {
	//
}

func transaction() {
	//
}
