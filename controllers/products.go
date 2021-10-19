package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mohamadakbar/go-invitee/models"
	"github.com/mohamadakbar/go-invitee/utils"
	"net/http"
	"time"
)

type Products struct {
	DB *gorm.DB
}

func NewProdsController (db *gorm.DB) *Products{
	return &Products{
		//repo: division.NewPSQLDivision(db),
		DB: db,
	}
}

func (cc *Products) GetAll(c *gin.Context)  {
	var products []models.Products
	//db := c.MustGet("db").(*gorm.DB)
	cc.DB.Find(&products)
	//db.Find(&products)

	res := utils.Response(http.StatusOK, "Success", products)
	c.JSON(http.StatusOK, res)
	return
}

func (cc *Products) Store(c *gin.Context)  {
	var req models.Products

	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, "error : "+err.Error())
		return
	}

	product := models.Products{
		Code:      req.Code,
		Name:      req.Name,
		Price:     req.Price,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	cc.DB.Create(&product)
	c.JSON(http.StatusOK, "Success insert data")
}