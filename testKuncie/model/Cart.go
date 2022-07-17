package model

import (
	"github.com/graphql-go/graphql"
)

type Cart struct {
	CartId     string   `json:"cart_id"`
	Item       []Items  `json:"item"`
	TotalPrice float64  `json:"total_price"`
	FreeGood   FreeItem `json:"free_good"`
}

type Items struct {
	SKU      string  `json:"sku"`
	Name     string  `json:"name"`
	QtyOrder int     `json:"qty_order"`
	Price    float64 `json:"price"`
}

type FreeItem struct {
	SKU  string `json:"sku"`
	Name string `json:"name"`
	Qty  int    `json:"qty"`
}

type AddCart struct {
	AddToCart Cart `json:"addtocart"`
}


var FreeItemType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "free_item",
		Fields: graphql.Fields{
			"sku": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var ItemType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "item",
		Fields: graphql.Fields{
			"sku": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"qty_order": &graphql.Field{
				Type: graphql.Int,
			},
			"price": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)

var CartType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "cart",
		Fields: graphql.Fields{
			"cart_id": &graphql.Field{
				Type: graphql.String,
			},
			"item": &graphql.Field{
				Type: graphql.NewList(ItemType),
			},
			"total_price": &graphql.Field{
				Type: graphql.Float,
			},
			"free_good": &graphql.Field{
				Type: FreeItemType,
			},
		},
	},
)
