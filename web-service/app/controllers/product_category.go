package controllers

import (
	"deeploy-exam/app/models"
	"deeploy-exam/app/repositorys"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductCategory struct{}

func NewProductCategoryController() *ProductCategory {
	return &ProductCategory{}
}

func (*ProductCategory) GetList(c *gin.Context) {
	var request models.SearchProductCategory
	c.ShouldBind(&request)

	var repo repositorys.Category
	list, count := repo.Getlist(&request)
	c.JSON(200, gin.H{
		"items":  list,
		"length": count,
	})
	return
}

func (*ProductCategory) GetInfo(c *gin.Context) {
	idStr := c.Param("category_id")
	categoryID, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "bad request body",
		})
		return
	}

	var repo repositorys.Category
	info := repo.GetInfo(uint(categoryID))
	c.JSON(200, gin.H{
		"item": info,
	})
	return
}

func (*ProductCategory) Create(c *gin.Context) {
	var request models.CreateProductCategory
	if c.BindJSON(&request) != nil {
		c.JSON(400, gin.H{
			"message": "bad request body",
		})
		return
	}

	var repo repositorys.Category
	if repo.Create(&request) != nil {
		c.JSON(500, gin.H{
			"message": "internal server error - cannot create category to database",
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "create completed",
		"item":    request,
	})
	return
}

func (*ProductCategory) Update(c *gin.Context) {
	var request models.UpdateProductCategory
	if c.BindJSON(&request) != nil {
		c.JSON(400, gin.H{
			"message": "bad request body",
		})
		return
	}

	var repo repositorys.Category
	if repo.Update(&request) != nil {
		c.JSON(500, gin.H{
			"message": "internal server error - cannot create category to database",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "update completed",
		"item":    request,
	})
	return
}

func (*ProductCategory) Delete(c *gin.Context) {
	idStr := c.Param("category_id")
	categoryID, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "bad request body",
		})
		return
	}

	var repo repositorys.Category
	if repo.Delete(uint(categoryID)) != nil {
		c.JSON(500, gin.H{
			"message": "internal server error - cannot create category to database",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "delete completed",
	})
	return
}
