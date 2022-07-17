package schema

import (
	"github.com/graphql-go/graphql"
	"kunciee/testKuncie/model"
	"kunciee/testKuncie/resolver"
	"log"

)

func CreateCartSchema() graphql.Schema {
	fields := graphql.Fields{
		"cart": &graphql.Field{
			Type:        model.CartType,
			Description: "Get cart By cart_id",
			Args: graphql.FieldConfigArgument{
				"cart_id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: resolver.GetCartById,
		},
		"list": &graphql.Field{
			Type:        graphql.NewList(model.CartType),
			Description: "Get Product List",
			Resolve:     resolver.GetListCart,
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery), Mutation: mutationCartType}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	return schema
}

var mutationCartType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"addtocart": &graphql.Field{
			//todo tipe diganti dengan model message success
			Type:        model.CartType,
			Description: "add product to cart",
			Args: graphql.FieldConfigArgument{
				"cart_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"item": &graphql.ArgumentConfig{
					Type: graphql.NewList(items),
				},
			},
			Resolve: resolver.AddProductToCart,
		},
	},
})

var items = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "items",
	Fields: graphql.InputObjectConfigFieldMap{
		"sku": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"qty_order": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})
