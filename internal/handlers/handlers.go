package handlers

import (
	"net/http"

)

func Handlers() {
	http.HandleFunc("/list/", getAllList)
	http.HandleFunc("/insert/", insertItem)
	http.HandleFunc("/remove/", removeItem)
	http.HandleFunc("/done/", updateItem)
}

func getAllList(w http.ResponseWriter, r *http.Request) {

}

func insertItem(w http.ResponseWriter, r *http.Request) {

}

func removeItem(w http.ResponseWriter, r *http.Request) {

}

func updateItem(w http.ResponseWriter, r *http.Request) {

}