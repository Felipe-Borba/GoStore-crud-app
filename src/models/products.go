package models

import (
	"simpleCrud/db"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func FindAll() []Product {
	database := db.Dbconnect()
	defer database.Close() //defer means: last thing to be executed

	findall, err := database.Query("select * from products order by id asc")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	productDAO := []Product{}

	for findall.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = findall.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity
		productDAO = append(productDAO, p)
	}

	return productDAO
}

func CreateNew(name, description string, price float64, quantity int) {
	db := db.Dbconnect()
	defer db.Close()

	insertData, err := db.Prepare("insert into products(product, description, price, quantity) values($1, $2, $3, $4)")
	if err != nil {
		//TODO in case of conversion error
		panic(err.Error())
	}

	insertData.Exec(name, description, price, quantity)
}

func Delete(id string) {
	db := db.Dbconnect()
	defer db.Close()

	deleteProduct, err := db.Prepare("delete from products where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deleteProduct.Exec(id)
}

func FindById(id string) Product {
	db := db.Dbconnect()
	defer db.Close()

	query, err := db.Query("select * from products where id=$1", id)
	if err != nil {
		//TODO in case of conversion error
		panic(err.Error())
	}

	productDAO := Product{}

	for query.Next() {
		err = query.Scan(&productDAO.Id, &productDAO.Name, &productDAO.Description, &productDAO.Price, &productDAO.Quantity)

		if err != nil {
			//TODO in case of conversion error
			panic(err.Error())
		}
	}

	return productDAO
}

func Update(id int, name, description string, price float64, quantity int) {
	db := db.Dbconnect()
	defer db.Close()

	query, err := db.Prepare("update products set product=$1, description=$2, price=$3, quantity=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}

	query.Exec(name, description, price, quantity, id)
}
