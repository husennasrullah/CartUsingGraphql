package resolver

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
	"kunciee/model"
)

//declare an empty array model, used for save data of cart
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
	return nil, fmt.Errorf("unknown data with this Cart_id")
}

func AddProductToCart(p graphql.ResolveParams) (result interface{}, err error) {
	var tempCart model.Cart
	for key, value := range p.Args {
		switch key {
		//case "cart_id":
		//	tempCart.CartId = value.(string)
		case "item":
			data := value.([]interface{})
			if len(data) == 0 {
				err = fmt.Errorf("item Cannot Empty")
				return
			}
			for i := 0; i < len(data); i++ {
				var temp model.Items
				byteData, _ := json.Marshal(data[i])
				_ = json.Unmarshal(byteData, &temp)

				if temp.SKU == "" || temp.QtyOrder == 0 {
					err = fmt.Errorf("SKU and Qty_Order Cannot Empty")
					return
				}

				tempCart.Item = append(tempCart.Item, temp)
			}
		}

		isOutOfStock, fieldErr, totalPrice := calculateOrder(&tempCart)
		if isOutOfStock {
			err = fmt.Errorf("item -> %v , Is Out Of Stock", fieldErr)
			return
		}

		tempCart.CartId = uuid.NewString()
		tempCart.TotalPrice = totalPrice
	}

	*CartData = append(*CartData, tempCart)
	result = tempCart
	return
}

func calculateOrder(orders *model.Cart) (status bool, fieldName string, totalPrice float64) {
	//get data product in existing ProductData memory
	product := *ProductData

	//loop over order item
	for i := 0; i < len(orders.Item); i++ {
		for j := 0; j < len(product); j++ {
			if product[j].SKU == orders.Item[i].SKU {

				if orders.Item[i].QtyOrder > product[j].Qty {
					fieldName = product[j].Name
					status = true
					return
				}

				price := CalculatePromotion(product[j], orders.Item[i].QtyOrder)
				totalPrice += price

				//set new qty to data product
				product[j].Qty = product[j].Qty - orders.Item[i].QtyOrder
				orders.Item[i].Name = product[j].Name
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
