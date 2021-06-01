package controllers

import (
	"net/http"
	"simpleCrud/models"
	"strconv"
	"text/template"
)

var htmlFiles = template.Must(template.ParseGlob("templates/*.html"))
var movedPermanently = 301

func Index(w http.ResponseWriter, r *http.Request) {
	htmlFiles.ExecuteTemplate(w, "Index", models.FindAll())
}

func New(w http.ResponseWriter, r *http.Request) {
	htmlFiles.ExecuteTemplate(w, "AddProduct", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		convertedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			//TODO in case conversion error
			panic(err.Error())
		}

		convertedQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			//TODO in case conversion error
			panic(err.Error())
		}

		models.CreateNew(name, description, convertedPrice, convertedQuantity)
	}
	http.Redirect(w, r, "/", movedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	models.Delete(productId)

	http.Redirect(w, r, "/", movedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	getID := r.URL.Query().Get("id")
	product := models.FindById(getID)
	htmlFiles.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		stringId := r.FormValue("id")
		id, err := strconv.Atoi(stringId)
		if err != nil {
			//TODO in case conversion error
			panic(err.Error())
		}
		name := r.FormValue("name")
		description := r.FormValue("description")
		stringPrice := r.FormValue("price")
		price, err := strconv.ParseFloat(stringPrice, 64)
		if err != nil {
			//TODO in case conversion error
			panic(err.Error())
		}
		stringQuantity := r.FormValue("quantity")
		quantity, err := strconv.Atoi(stringQuantity)
		if err != nil {
			//TODO in case conversion error
			panic(err.Error())
		}

		models.Update(id, name, description, price, quantity)
	}
	http.Redirect(w, r, "/", movedPermanently)
}
