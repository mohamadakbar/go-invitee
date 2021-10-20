package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mohamadakbar/go-invitee/repo"
	"github.com/mohamadakbar/go-invitee/repo/paket"
	"github.com/mohamadakbar/go-invitee/requests"
	"github.com/mohamadakbar/go-invitee/utils"
	"net/http"
)

type Pakets struct {
	DB   *gorm.DB
	repo repo.Paket
}

func NewpaketController(db *gorm.DB) *Pakets {
	return &Pakets{
		DB:   db,
		repo: paket.NewSQLPaketRepo(db),
	}
}

func (cc *Pakets) GetPaketById(c *gin.Context) {
	var req requests.FindPaket

	err := c.ShouldBindJSON(&req)
	if err != nil {
		res := utils.Response(http.StatusBadRequest, "error : "+err.Error(), nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	paket, sts, err := cc.repo.InqGetById(req.ID)
	if err != nil {
		res := utils.Response(sts, err.Error(), nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.Response(http.StatusOK, "success", paket)
	c.JSON(http.StatusOK, res)
	return
}
