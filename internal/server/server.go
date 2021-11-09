package server

import (
	"net/http"
	
	"github.com/mthsrocha/mb/internal/handlers"
)

type Server struct {

}

func HttpServer() {
	handlers.Handlers()
	http.ListenAndServe(":5012", nil)
}