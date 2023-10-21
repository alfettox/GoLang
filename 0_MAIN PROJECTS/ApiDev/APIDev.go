/*
Author: Giovanni De Franceschi
*/






package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux" //Router handler
)

type Account struct {
	Number  string `json:"AccountNumber"`
	Balance string `json:"Balance"`
	Desc    string `json:"AccountDescription"`
}

var Accounts []Account

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllAccounts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllAccounts")
	json.NewEncoder(w).Encode(Accounts)
}

func returnAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // Corrected function name
	key := vars["number"]
	for _, account := range Accounts {
		if account.Number == key {
			json.NewEncoder(w).Encode(account)
			return
		}
	}
	// If the account with the specified number is not found
	http.NotFound(w, r)
}

func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/accounts", returnAllAccounts).Methods("GET")
	router.HandleFunc("/account/{number}", returnAccount).Methods("GET")
	router.HandleFunc("/account", createAccount).Methods("POST")
	router.HandleFunc("/account/{number}", createAccount).Methods("PUT")
	router.HandleFunc("/account/{number}", deleteAccount).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":10000", router))
}

func main() {
	Accounts = []Account{
		Account{Number: "1001", Balance: "1000", Desc: "Savings"},
		Account{Number: "1002", Balance: "2000", Desc: "Checking"},
	}
	handleRequests()
}

func createAccount(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var account Account
	json.Unmarshal(requestBody, &account)
	Accounts = append(Accounts, account)
	json.NewEncoder(w).Encode(account)
}

func deleteAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["number"]

	for index, account := range Accounts {
		if account.Number == id {
			Accounts = append(Accounts[:index], Accounts[index+1:]...) //... means all the elements
		}
	}
}

func updateAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["number"]

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var updatedAccount Account
	err = json.Unmarshal(requestBody, &updatedAccount) // Corrected from err.json.Unmarshal to err = json.Unmarshal
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for index, account := range Accounts {
		if account.Number == id {
			Accounts[index] = updatedAccount
			json.NewEncoder(w).Encode(updatedAccount)
			return
		}
	}

	http.NotFound(w, r)
}
