package main

import (
	"database/sql"
	"html/template"
	"net/http"
	_ "github.com/lib/pq"
)

type Products struct {
	Id			int
	Name		string
	Description	string
	Price 		float64
	Quantity	int
} 

var readTemplates = template.Must(template.ParseGlob("templates/*.html")) //le todos templates que tem .html


func connPostgres() *sql.DB {
	connection := "user=postgres dbname=golang-marketplace password=admin host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err)
	}
	return db

}

func main() {
	http.HandleFunc("/", index)   //toda vez que chamar a barra, executo a func index
	http.ListenAndServe(":8000", nil)
}


func index(w http.ResponseWriter, r *http.Request){  //(w writer, r responser)
	
	db := connPostgres()
	selectProducts, err :=db.Query("select * from tbproducts")
	if err != nil {
		panic(err.Error())
	}

	eachProduct := Products{}
	allProducts := []Products{}
	for selectProducts.Next(){
		var id, quantity int
		var productname, description string
		var price float64

		err = selectProducts.Scan(&id, &productname, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		eachProduct.Name = productname
		eachProduct.Description = description
		eachProduct.Price = price
		eachProduct.Quantity = quantity

		allProducts = append(allProducts, eachProduct)
	}

	readTemplates.ExecuteTemplate(w, "index", allProducts)
	defer db.Close()
}
