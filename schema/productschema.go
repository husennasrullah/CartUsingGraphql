package schema

import (
	"github.com/graphql-go/graphql"
	"kunciee/model"
	"kunciee/resolver"
	"log"
)

var mutationProductType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"create": &graphql.Field{
			Type:        model.ProductType,
			Description: "Create a new Tutorial",
			Args: graphql.FieldConfigArgument{
				"sku": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"price": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Float),
				},
				"qty": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: resolver.CreateProduct,
		},
	},
})

func CreateProductSchema() graphql.Schema {
	fields := graphql.Fields{
		"product": &graphql.Field{
			Type:        model.ProductType,
			Description: "Get product By ID",
			Args: graphql.FieldConfigArgument{
				"sku": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"qty": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: resolver.GetProductById,
		},
		"list": &graphql.Field{
			Type:        graphql.NewList(model.ProductType),
			Description: "Get Product List",
			Resolve:     resolver.GetlistProduct,
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery), Mutation: mutationProductType}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	return schema
}

//func ResolveAddToCart (params graphql.ResolveParams) (interface{}, error) {
//	return
//}
