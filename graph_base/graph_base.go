package graph_base

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bervProject/go-microservice-boilerplate/models"
	"github.com/graphql-go/graphql"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var global_schema graphql.Schema

type PostData struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func InitGraphQL(db *gorm.DB) {
	// Schema
	userType := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "User",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
		"getAllUsers": &graphql.Field{
			Type: graphql.NewList(userType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var users []models.User
				db.Model(&models.User{}).Find(&users)
				return users, nil
			},
		},
	}
	mutationFields := graphql.Fields{
		"createUser": &graphql.Field{
			Type:        userType,
			Description: "Create New User",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				user := models.User{
					Name: params.Args["name"].(string),
				}
				result := db.Create(&user)
				if result.Error != nil {
					return nil, result.Error
				}
				return user, nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	rootMutation := graphql.ObjectConfig{Name: "RootMutation", Fields: mutationFields}
	schemaConfig := graphql.SchemaConfig{
		Query:    graphql.NewObject(rootQuery),
		Mutation: graphql.NewObject(rootMutation),
	}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	global_schema = schema
}

func GraphQLHandler(c echo.Context) error {
	data := new(PostData)
	if err := c.Bind(data); err != nil {
		return err
	}
	result := graphql.Do(graphql.Params{
		Schema:         global_schema,
		RequestString:  data.Query,
		VariableValues: data.Variables,
		OperationName:  data.Operation,
	})
	return c.JSON(http.StatusOK, result)
}
