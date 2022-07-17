package resolver

import (
	"github.com/graphql-go/graphql"
	"kunciee/model"
)

var InitiateData = model.PopulateProduct()
var ProductData = &InitiateData

func GetlistProduct(params graphql.ResolveParams) (interface{}, error) {
	//get data to db
	return ProductData, nil
}

func GetProductById(params graphql.ResolveParams) (interface{}, error) {
	for key, tes := range params.Args {
		switch key {
		case "sku":
			sku, ok := tes.(string)
			if ok {
				for _, data := range *ProductData {
					if data.SKU == sku {
						return data, nil
					}
				}
			}
		case "qty":
			qty, ok := tes.(int)
			if ok {
				for _, data := range *ProductData {
					if data.Qty == qty {
						return data, nil
					}
				}
			}
		}
	}
	return nil, nil
}

func CreateProduct(params graphql.ResolveParams) (interface{}, error) {
	var product model.Product
	product.SKU, _ = params.Args["sku"].(string)
	product.Name, _ = params.Args["name"].(string)
	product.Qty, _ = params.Args["qty"].(int)
	product.Price, _ = params.Args["price"].(float64)

	*ProductData = append(*ProductData, product)

	return product, nil
}
