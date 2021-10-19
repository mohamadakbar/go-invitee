package main

import (
	"fmt"
	_ "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	config "github.com/mohamadakbar/go-invitee/config"
	utils "github.com/mohamadakbar/go-invitee/utils"
)

func main() {
	port := utils.Env("PORT")
	var db = config.GetConn()
	defer db.Close()

	config := &config.Routes{
		DB: db,
	}

	routes := config.SetupRoutes()
	routes.Run(fmt.Sprintf(":%s", port))

}