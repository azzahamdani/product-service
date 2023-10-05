package main

import (
	"github.com/gin-gonic/gin"
	"github.com/azzahamdani/product-service/handler"
	"github.com/azzahamdani/product-service/repository"
	"github.com/azzahamdani/product-service/service"
)

func main() {
	repo := repository.NewProductRepository()
	svc := service.NewProductService(repo)
	productHandler := handler.NewProductHandler(svc)

	router := gin.Default()
	router.GET("/product/:productID", productHandler.GetProduct)
	router.Run(":5001")
}

