package controllers

import (
	"deeploy-exam/app/models"
	"deeploy-exam/app/repositorys"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct{}

func NewProductController() *Product {
	return &Product{}
}

func (*Product) GetList(c *gin.Context) {
	var request models.SearchProduct
	c.ShouldBind(&request)

	var repo repositorys.Product
	list, count := repo.Getlist(&request)
	c.JSON(200, gin.H{
		"items":  list,
		"length": count,
	})
	return
}

func (*Product) GetInfo(c *gin.Context) {
	idStr := c.Param("product_id")
	productID, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "bad request body",
		})
		return
	}

	var repo repositorys.Product
	info := repo.GetInfo(uint(productID))
	c.JSON(200, gin.H{
		"item": info,
	})
	return
}

func (*Product) Create(c *gin.Context) {
	var request models.CreateProduct
	if c.BindJSON(&request) != nil {
		c.JSON(400, gin.H{
			"message": "bad request body",
		})
		return
	}

	var repo repositorys.Product
	if repo.Create(&request) != nil {
		c.JSON(500, gin.H{
			"message": "internal server error - cannot create product to database",
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "create completed",
		"item":    request,
	})
	return
}

func (*Product) Update(c *gin.Context) {
	var request models.UpdateProduct
	if c.BindJSON(&request) != nil {
		c.JSON(400, gin.H{
			"message": "bad request body",
		})
		return
	}

	var repo repositorys.Product
	if repo.Update(&request) != nil {
		c.JSON(500, gin.H{
			"message": "internal server error - cannot create product to database",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "update product completed",
		"item":    request,
	})
	return
}

func (*Product) Delete(c *gin.Context) {
	idStr := c.Param("product_id")
	productID, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "bad request body",
		})
		return
	}

	var repo repositorys.Product
	if repo.Delete(uint(productID)) != nil {
		c.JSON(500, gin.H{
			"message": "internal server error - cannot create product to database",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "delete completed",
	})
	return
}
