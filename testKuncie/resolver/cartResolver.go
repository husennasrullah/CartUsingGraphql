package resolver

import (
	"encoding/json"
	"github.com/graphql-go/graphql"
	"kunciee/testKuncie/model"
)

var InitiateCart []model.Cart
var CartData = &InitiateCart

func GetListCart(params graphql.ResolveParams) (interface{}, error) {
	//get data to db
	return CartData, nil
}

func GetCartById(p graphql.ResolveParams) (interface{}, error) {
	for key, value := range p.Args {
		switch key {
		case "cart_id":
			id, ok := value.(string)
			if ok {
				for _, data := range *CartData {
					if data.CartId == id {
						return data, nil
					}
				}
			}
		}
	}
	return nil, nil
}

func AddProductToCart(p graphql.ResolveParams) (interface{}, error) {
	var tempCart model.Cart
	for key, value := range p.Args {
		switch key {
		case "cart_id":
			tempCart.CartId = value.(string)
		case "item":
			data := value.([]interface{})
			if len(data) == 0 {
				//todo return error
			}
			for i := 0; i < len(data); i++ {
				var temp model.Items
				byteData, _ := json.Marshal(data[i])
				_ = json.Unmarshal(byteData, &temp)

				if temp.SKU == "" || temp.QtyOrder == 0 {
					//todo return error
				}

				tempCart.Item = append(tempCart.Item, temp)
			}
		}

		isReadyStock, totalPrice := calculateOrder(&tempCart)
		if !isReadyStock {
			//todo return error
		}

		tempCart.TotalPrice = totalPrice

	}
	*CartData = append(*CartData, tempCart)
	return tempCart, nil
}

func calculateOrder(orders *model.Cart) (status bool, totalPrice float64) {

	//get data product
	product := *ProductData

	//loop over order item
	for _, order := range orders.Item {
		for i := 0; i < len(product); i++ {
			if product[i].SKU == order.SKU {
				if order.QtyOrder > product[i].Qty {
					status = false
					return
				}

				price := CalculatePromotion(product[i], order.QtyOrder)
				totalPrice += price

				//set new qty to data product
				product[i].Qty = product[i].Qty - order.QtyOrder
			}
		}
	}

	return
}

func CalculatePromotion(product model.Product, qty int) (price float64) {
	switch product.SKU {
	//google home
	case "120P90":
		price2 := 2 * product.Price
		// buy 3 google home price of 2
		if qty > 2 {
			if qty%3 == 0 {
				price = float64(qty/3) * price2
			} else {
				price = (float64(rune(qty/3)) * price2) + (float64(qty%3) * product.Price)
			}
		} else {
			price = float64(qty) * product.Price
		}

	case "43N23P": //macbook pro
		// tiap pembelian 1 dapat bonus raspberry pi
		price = float64(qty) * product.Price
		//kurangi qty rapberry pi untuk pembelian ini
		productDataTemp := *ProductData
		for i := 0; i < len(productDataTemp); i++ {
			if productDataTemp[i].SKU == "234234" {
				if productDataTemp[i].Qty > 1 {
					productDataTemp[i].Qty -= 1
				}
			}
		}

	case "A304SD": //alexa speaker
		// beli lebih dari 3 alexa speaker discount 10%
		if qty > 2 {
			price = (float64(qty) * product.Price) - (float64(qty) * product.Price * 10 / 100)
		} else {
			price = float64(qty) * product.Price
		}

	case "234234": //raspberry pi
		price = float64(qty) * product.Price
	}

	return
}
