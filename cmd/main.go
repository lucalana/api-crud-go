package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	// Camada de Repository
	ProductRepository := repository.NewProdutcRepository(dbConnection)
	// Camada useCase
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)

	// Camada de controller
	productController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", productController.GetProducts)
	server.POST("/products", productController.CreateProduct)
	server.GET("/product/:id", productController.GetProductById)

	server.Run(":8000")
}
