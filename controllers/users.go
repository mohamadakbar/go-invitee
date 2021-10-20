package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mohamadakbar/go-invitee/models"
	"github.com/mohamadakbar/go-invitee/repo"
	"github.com/mohamadakbar/go-invitee/repo/paket"
	"github.com/mohamadakbar/go-invitee/repo/user"
	"github.com/mohamadakbar/go-invitee/requests"
	"github.com/mohamadakbar/go-invitee/utils"
	"net/http"
)

type Users struct {
	DB        *gorm.DB
	repo      repo.User
	repoPaket repo.Paket
}

func NewusersController(db *gorm.DB) *Users {
	return &Users{
		DB:        db,
		repo:      user.NewSQLUserRepo(db),
		repoPaket: paket.NewSQLPaketRepo(db),
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
	id := c.Param("slug")

	//fmt.Println(id)
	//os.Exit(1)

	err := cc.DB.Where("slug = ? ", id).First(&user).Error
	if err != nil {
		res := utils.Response(http.StatusNotFound, "cannot find record for id "+id, nil)
		c.JSON(http.StatusNotFound, res)
		return
	}
	res := utils.Response(http.StatusOK, "Success", user)
	c.JSON(http.StatusOK, res)
	return
}

func (cc *Users) GetBySlug(c *gin.Context) {
	var req requests.StoreUser

	err := c.ShouldBindJSON(&req)
	if err != nil {
		res := utils.Response(http.StatusBadRequest, "error : "+err.Error(), nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	users, sts, err := cc.repo.InqGetBySlug(req.Slug)
	if err != nil {
		res := utils.Response(sts, err.Error(), nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.Response(http.StatusOK, "success", users)
	c.JSON(http.StatusOK, res)
	return
}

func (cc *Users) Store(c *gin.Context) {
	var req requests.StoreUser

	err := c.ShouldBindJSON(&req)
	if err != nil {
		res := utils.Response(http.StatusBadRequest, "error : "+err.Error(), nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	// TODO: start validasi
	_, stsSlug, _ := cc.repo.InqGetBySlug(req.Slug)
	if stsSlug == 200 {
		res := utils.Response(http.StatusForbidden, "slug already exist", nil)
		c.JSON(http.StatusForbidden, res)
		return
	}

	_, stsEmail, _ := cc.repo.InqGetByEmail(req.Email)
	if stsEmail == 200 {
		res := utils.Response(http.StatusForbidden, "email already exist", nil)
		c.JSON(http.StatusForbidden, res)
		return
	}

	paket, sts, err := cc.repoPaket.InqGetById(req.IDPaket)

	// TODO: start store user
	users, sts, err := cc.repo.InqStore(req, paket)
	if err != nil {
		res := utils.Response(sts, err.Error(), nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.Response(http.StatusOK, "success", users)
	c.JSON(http.StatusOK, res)
	return
}
