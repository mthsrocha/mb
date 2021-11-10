package main

import (
	"fmt"

	"github.com/mthsrocha/mb/internal/server"
)


func main() {

	fmt.Println("Starting web application")


	server.HttpServer()
}
