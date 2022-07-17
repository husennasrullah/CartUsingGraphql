package main

import (
	gqlhandler "github.com/graphql-go/graphql-go-handler"
	"kunciee/testKuncie/model"
	"kunciee/testKuncie/schema"
	"log"
	"net/http"


)

var InitiateData = model.PopulateProduct()
var DataInventory = &InitiateData

func main() {
	productSchema := schema.CreateProductSchema()
	cartSchema := schema.CreateCartSchema()
	//test

	handlerproduct := gqlhandler.New(&gqlhandler.Config{Schema: &productSchema})
	handlerCart := gqlhandler.New(&gqlhandler.Config{Schema: &cartSchema})

	http.Handle("/product", handlerproduct)
	http.Handle("/cart", handlerCart)

	log.Println("Server started at http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))

}
