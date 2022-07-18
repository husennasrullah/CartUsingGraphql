package unittest

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/stretchr/testify/assert"
	"kunciee/model"
	"kunciee/schema"
	"testing"
)

func TestCart(t *testing.T) {
	tests := []struct {
		name          string
		query         string
		expectedPrice float64
	}{
		{
			name:          "Scanned Items: Google Home, Google Home, Google Home",
			query:         "mutation {addtocart(cart_id: \"test1\"item: [{sku: \"120P90\"qty_order : 3}]){cart_id item { sku name qty_order }total_price free_good {sku name}}}",
			expectedPrice: 99.98,
		},
		{
			name:          "Scanned Items: Alexa Speaker, Alexa Speaker, Alexa Speaker",
			query:         "mutation {addtocart(cart_id: \"test2\"item: [{sku: \"A304SD\"qty_order : 3}]){cart_id item { sku name qty_order }total_price free_good {sku name}}}",
			expectedPrice: 295.65,
		},
		{
			name:          "Scanned Items: MacBook Pro, Raspberry Pi B",
			query:         "mutation {addtocart(cart_id: \"test3\"item: [{sku: \"43N23P\"qty_order : 1}, {sku: \"234234\"qty_order : 1}]){cart_id item { sku name qty_order }total_price free_good {sku name}}}",
			expectedPrice: 5399.99,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cartSchema := schema.CreateCartSchema()
			params := graphql.Params{Schema: cartSchema, RequestString: test.query}
			resp := graphql.Do(params)

			if len(resp.Errors) > 0 {
				assert.FailNow(t, "failed to execute graphql operation, errors: %+v", resp.Errors)
			}

			var data model.AddCart
			jsonStr, _ := json.Marshal(resp.Data)
			_ = json.Unmarshal(jsonStr, &data)

			assert.Equal(t, test.expectedPrice, data.AddToCart.TotalPrice, fmt.Sprintf("Price Must Be -->  %v", test.expectedPrice))
		})
	}
}
