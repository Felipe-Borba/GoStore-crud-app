package main

import (
	"net/http"
	"simpleCrud/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8080", nil)
}
