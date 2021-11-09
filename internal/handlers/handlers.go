package handlers

import (
	"net/http"

	"github.com/mthsrocha/mb/main"
)

func Handlers() {
	http.HandleFunc("/list", main.GetAllList)
	http.HandleFunc("/insert", main.GetAllList)
	http.HandleFunc("/remove", main.GetAllList)
	http.HandleFunc("/done", main.GetAllList)
}
