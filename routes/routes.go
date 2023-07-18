package routes

import (
	"net/http"
	"remote-repo.com/lucas/webapp/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index) //toda vez que chamar a barra, executo a func index
	http.HandleFunc("/new", controllers.NewProductPage)
	http.HandleFunc("/save-product", controllers.InsertProduct)
	http.HandleFunc("/delete", controllers.DeleteProduct)
}
