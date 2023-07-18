package main

import (
	"net/http"
	"remote-repo.com/lucas/webapp/routes"
)



func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}

