package main

import (
	"food-helper/config"
	"food-helper/dao"
	"food-helper/services"
	"food-helper/web"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		panic(err)
	}
	dishDAO := dao.NewDishDAO(db)
	dishService := services.NewDishService(dishDAO)

	server := gin.Default()

	// 添加静态文件服务
	server.LoadHTMLGlob("templates/*")
	server.Static("/static", "./static")

	// 首页路由
	server.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	dishHandler := web.NewDishHandler(dishService)
	dishHandler.RegisterRoutes(server)
	server.Run(":8080")
}
