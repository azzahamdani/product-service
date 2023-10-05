package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/azzahamdani/product-service/service"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	productID := c.Param("productID")
	product, err := h.service.GetProduct(productID)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, product)
}
