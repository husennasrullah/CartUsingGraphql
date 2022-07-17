package model

import (
	"github.com/graphql-go/graphql"
)

type Product struct {
	SKU   string  `json:"sku"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Qty   int     `json:"qty"`
}

var ProductType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Product",
		Fields: graphql.Fields{
			"sku": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"price": &graphql.Field{
				Type: graphql.Float,
			},
			"qty": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

func PopulateProduct() []Product {
	var products []Product

	products = []Product{
		{
			SKU:   "120P90",
			Name:  "Google Home",
			Price: 49.99,
			Qty:   10,
		},
		{
			SKU:   "43N23P",
			Name:  "MacBook Pro",
			Price: 5399.99,
			Qty:   5,
		},
		{
			SKU:   "A304SD",
			Name:  "Alexa Speaker",
			Price: 109.50,
			Qty:   10,
		},
		{
			SKU:   "234234",
			Name:  "Raspberry Pi",
			Price: 30.00,
			Qty:   2,
		},
	}

	return products
}
