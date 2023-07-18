package models

import "remote-repo.com/lucas/webapp/db"

type Products struct {
	Id			int
	Name		string
	Description	string
	Price 		float64
	Quantity	int
} 


func SearchAllProducts() []Products {
		
	db := db.ConnPostgres()
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
	defer db.Close()
	return allProducts
}

func SaveNewProduct(productname string, description string, price float64, quantity int) {
	db := db.ConnPostgres()
	insertNewProduct, err := db.Prepare("insert into tbproducts(productname,description,price,quantity) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insertNewProduct.Exec(productname, description, price, quantity)
	defer db.Close()
}