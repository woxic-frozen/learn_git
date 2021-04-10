package main

import (
	"ginblog/function"
	"github.com/gin-gonic/gin"
)
func main() {
	router:=gin.Default()

	router.POST("/login",function.Login)
	router.POST("/register",function.Register)
	usergroup:=router.Group("/user")
	usergroup.Use(function.MiddleWare())
	{
		usergroup.POST("/launchArticle",function.Article)
		usergroup.POST("/message",function.Rreview)
		usergroup.GET("/like/:aid",function.Likes)
		usergroup.POST("/article",function.OneArticle)
	}
    //router.POST("/user")
	router.Run(":8080")
}