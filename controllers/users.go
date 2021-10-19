package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mohamadakbar/go-invitee/models"
	"github.com/mohamadakbar/go-invitee/utils"
	"net/http"
)

type Users struct {
	DB *gorm.DB
}

func NewusersController (db *gorm.DB) *Users{
	return &Users{
		//repo: division.NewPSQLDivision(db),
		DB: db,
	}
}

func (cc *Users) GetAll(c *gin.Context) {
	var users []models.Users
	cc.DB.Find(&users)

	res := utils.Response(http.StatusOK, "Success", users)
	c.JSON(http.StatusOK, res)
	return
}

func (cc *Users) GetById(c *gin.Context) {
	var user models.Users
	id := c.Param("id")

	err := cc.DB.Where("id = ? ", id).First(&user).Error
	if err != nil {
		res := utils.Response(http.StatusNotFound, "cannot find record id "+id, nil)
		c.JSON(http.StatusNotFound, res)
		return
	}
	res := utils.Response(http.StatusOK, "Success", user)
	c.JSON(http.StatusOK, res)
	return
}