package controllers

import (
	"deeploy-exam/app/models"
	"deeploy-exam/app/repositorys"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StoreProduct struct{}

func NewStoreProductController() *StoreProduct {
	return &StoreProduct{}
}

func (*StoreProduct) Create(c *gin.Context) {
	var request []models.CreateStoreProduct
	if c.BindJSON(&request) != nil {
		c.JSON(400, gin.H{
			"message": "bad request body",
		})
		return
	}

	var spRepo repositorys.StoreProduct
	if spRepo.Create(request) != nil {
		c.JSON(500, gin.H{
			"message": "internal server error - cannot create store's product to database",
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "create completed",
		"item":    request,
	})
	return
}

func (*StoreProduct) Update(c *gin.Context) {
	var request models.UpdateStoreProduct
	if c.BindJSON(&request) != nil {
		c.JSON(400, gin.H{
			"message": "bad request body",
		})
		return
	}

	var repo repositorys.StoreProduct
	if repo.Update(&request) != nil {
		c.JSON(500, gin.H{
			"message": "internal server error - cannot update store's product to database",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "update completed",
		"item":    request,
	})
	return
}

func (*StoreProduct) Delete(c *gin.Context) {
	idStr := c.Param("req_id")
	reqID, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "bad request body",
		})
		return
	}

	var repo repositorys.StoreProduct
	if repo.Delete(uint(reqID)) != nil {
		c.JSON(500, gin.H{
			"message": "internal server error - cannot delete store's product from database",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "delete completed",
	})
	return
}
