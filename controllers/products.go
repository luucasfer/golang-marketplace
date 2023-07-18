package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"remote-repo.com/lucas/webapp/models"
)

var readTemplates = template.Must(template.ParseGlob("templates/*.html")) //le todos templates que tem .html

func Index(w http.ResponseWriter, r *http.Request) { //(w writer, r responser)
	allProducts := models.SearchAllProducts()
	readTemplates.ExecuteTemplate(w, "index", allProducts)
}

func NewProductPage(w http.ResponseWriter, r *http.Request) {
	readTemplates.ExecuteTemplate(w, "NewProducts", nil)
}

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name 		:= r.FormValue("name")
		description := r.FormValue("description")
		price 		:= r.FormValue("price")
		quantity 	:= r.FormValue("quantity")

		priceFloat, err 	:= strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error in price convertion")
		}

		quantityInt, err 	:= strconv.Atoi(quantity)
		if err != nil {
			log.Println("Error in quantity convertion")
		}

		models.SaveNewProduct(name, description, priceFloat, quantityInt)
	}
	http.Redirect(w, r, "/", 301)
}