package config

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mohamadakbar/go-invitee/controllers"
)

type Routes struct {
	DB *gorm.DB
}

func (r *Routes) SetupRoutes() *gin.Engine {

	app := gin.Default()
	main := app.Group("/")
	app.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	//app.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//app.GET("/products", Prods.GetAll)

	prodCtrl := controllers.NewProdsController(r.DB)
	products := main.Group("/")
	{
		products.GET("products", prodCtrl.GetAll)
		products.POST("product", prodCtrl.Store)
	}

	userCtrl := controllers.NewusersController(r.DB)
	users := main.Group("/")
	{
		users.GET("users", userCtrl.GetAll)
		users.GET("user/:id", userCtrl.GetById)
		users.POST("user/slug", userCtrl.GetBySlug)
		users.POST("user", userCtrl.Store)
	}

	paketCtrl := controllers.NewpaketController(r.DB)
	paket := main.Group("/")
	{
		paket.POST("paket", paketCtrl.GetPaketById)
		//paket.GET("user/:id", paketCtrl.GetById)
		//paket.POST("user/slug", paketCtrl.GetBySlug)
		//paket.POST("user", paketCtrl.Store)
	}

	return app
}
