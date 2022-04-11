package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bervProject/go-microservice-boilerplate/graph_base"
	"github.com/bervProject/go-microservice-boilerplate/models"
	"github.com/bervProject/go-microservice-boilerplate/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := utils.GetEnv("DB_CONNECTION_STRING", "")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database, error: %v", err)
	}
	db.AutoMigrate(&models.User{})
	graph_base.InitGraphQL(db)
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/graphql", graph_base.GraphQLHandler)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", utils.GetEnv("PORT", "1323"))))
}
