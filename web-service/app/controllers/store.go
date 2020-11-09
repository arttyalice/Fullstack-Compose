package controllers

import (
	"deeploy-exam/app/models"
	"deeploy-exam/app/repositorys"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Store struct{}

func NewStoreController() *Store {
	return &Store{}
}

func (*Store) GetList(c *gin.Context) {
	var request models.SearchStore
	c.ShouldBind(&request)

	var repo repositorys.Store
	list, count := repo.Getlist(&request)
	c.JSON(200, gin.H{
		"items":  list,
		"length": count,
	})
	return
}

func (*Store) GetInfo(c *gin.Context) {
	idStr := c.Param("store_id")
	storeID, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "bad request body",
		})
		return
	}

	var repo repositorys.Store
	info := repo.GetInfo(uint(storeID))
	c.JSON(200, gin.H{
		"item": info,
	})
	return
}

func (*Store) Create(c *gin.Context) {
	var request models.CreateStore
	if c.BindJSON(&request) != nil {
		c.JSON(400, gin.H{
			"message": "bad request body",
		})
		return
	}

	var repo repositorys.Store
	if repo.Create(&request) != nil {
		c.JSON(500, gin.H{
			"message": "internal server error - cannot create store to database",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "create completed",
		"item":    request,
	})
	return
}

func (*Store) Update(c *gin.Context) {
	var request models.UpdateStore
	if c.BindJSON(&request) != nil {
		c.JSON(400, gin.H{
			"message": "bad request body",
		})
		return
	}

	var repo repositorys.Store
	if repo.Update(&request) != nil {
		c.JSON(500, gin.H{
			"message": "internal server error - cannot create store to database",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "update completed",
		"item":    request,
	})
	return
}

func (*Store) Delete(c *gin.Context) {
	idStr := c.Param("store_id")
	storeID, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "bad request body",
		})
		return
	}

	var repo repositorys.Store
	if repo.Delete(uint(storeID)) != nil {
		c.JSON(500, gin.H{
			"message": "internal server error - cannot create store to database",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "delete completed",
	})
	return
}
