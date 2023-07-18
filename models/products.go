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
	selectProducts, err :=db.Query("select * from tbproducts order by id asc")
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

		eachProduct.Id = id
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

func DeleteSelectedProduct(id string) {
	db := db.ConnPostgres()
	delete, err := db.Prepare("delete from tbproducts where id = $1")
	if err != nil {
		panic(err.Error())
	}

	delete.Exec(id)
	defer db.Close()
}

func EditSelectedProduct(id string) Products {
	db := db.ConnPostgres()
	productToEdit, err := db.Query("select * from tbproducts where id = $1", id)
	if err != nil {
		panic(err.Error())
	}

	newProductData := Products{}
	for productToEdit.Next(){
		var id, quantity int
		var productname, description string
		var price float64

		err = productToEdit.Scan(&id, &productname, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		newProductData.Id = id
		newProductData.Name = productname
		newProductData.Description = description
		newProductData.Price = price
		newProductData.Quantity = quantity 
	}
	defer db.Close()
	return newProductData
}

func UpdateSelectedProduct(id int, productname string, description string, price float64, quantity int){
	db := db.ConnPostgres()
	productToUpdate, err := db.Prepare("update tbproducts set productname=$2, description=$3, price=$4, quantity=$5 where id = $1")
	if err != nil {
		panic(err.Error())
	}

	productToUpdate.Exec(id, productname, description, price, quantity)
	defer db.Close()
}