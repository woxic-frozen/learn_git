package function

import (
	_ "database/sql"
	"ginblog/sql1"
	strcutinf "ginblog/structinf1"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)
//登录
func Login(c *gin.Context){
	var user strcutinf.User//用户结构体在structinf文件中
	err := c.ShouldBind(&user)
	if err != nil {
		log.Fatal(err)
	}
	//find函数为查找数据库中是否有对应函数在sql文件中
	if sql.Find(user.Id,user.Password){
       c.SetCookie("id","123",1200,"/","localhost",false,true)//中间件使用
       c.SetCookie("uid",user.Id,1200,"/","localhost",false,true)//cookie储存用户id
		c.JSON(200,gin.H{
			"status":"ok",
			"message":"login succeed",
		})
	}else{
		c.JSON(200,gin.H{
			"status":"ok",
			"message":"login failed id or password wrong",
		})
	}
}
// 注册函数
func Register(c *gin.Context){
	var user strcutinf.User//用户结构体在structinf文件中
	err:=c.Bind(&user)
	if err!=nil{
		panic(err)
	}
	//注册sql操作判断是否可以注册如可以则注册
	if sql.Register(user.Id,user.Password){
		c.JSON(200,gin.H{
			"status" :"ok",
			"message":"register succeed",
		})
	}else{
		c.JSON(200,gin.H{
			"status":"error",
			"message":"id has been used",
		})
	}
}
//发布文章
func Article(c *gin.Context){
      var article strcutinf.ArticleInfo//文章结构体在structinf中
      err:=c.ShouldBind(&article)
      if err!=nil{
      	panic(err)
	  }
	  id, err := c.Cookie("uid")//从cookie中获取作者id
	  article.Id=id
	  if err!=nil{
	  }
      sql.Luancharticle(article)
	  c.JSON(200,gin.H{
	  	"message":"article launch succeed",
	  })
}
//发表文章评论
func Rreview(c *gin.Context){
	 var message strcutinf.Message//评论结构体在structinf中
	 aid:=c.Query("aid")//获取文章aid，aid为文章表的自增值
	 message.Aid,_=strconv.Atoi(aid)
	 err:=c.ShouldBind(&message)
	 id,err:=c.Cookie("uid")
	 if err!=nil{
	 	panic(err)
	 }
	 message.Id=id
	 if err!=nil{
      	panic(err)
	  }
	  sql.Luanchmessge(message)
	 c.JSON(200,gin.H{
	 	"message":"评论发表成功",
	 })


}
//对文章点赞
func Likes(c *gin.Context){
      aid:=c.Param("aid")//从url中获取文章aid确定文章
      aid1,err:=strconv.Atoi(aid)
      if err!=nil{
      	c.JSON(400,gin.H{
      		"message":"error",
		})
      	panic(err)
	  }
      sql.Likes(aid1)//sql操作在sql文件中
      c.JSON(200,gin.H{
      	"message":"give a like succeed",
	  })
}
//返回所有对应文章的文章内容和评论
func OneArticle(c *gin.Context){
	aid:=c.Query("aid")
	aidint,err:=strconv.Atoi(aid)
	if err!=nil{
		panic(err)
	}
	article,message:=sql.QueryArticle(aidint)
	c.JSON(200,gin.H{
		"status":"ok",
		"statusaid":aid,
		"article":article,
	})
   for _,v:=range message{
   	c.JSON(200,gin.H{
   		"status":"ok",
   		"message":v,
	})
   }

}
func MiddleWare()gin.HandlerFunc{
	return func(c*gin.Context){
		if cookie ,err:=c.Request.Cookie("id");err==nil{
			value := cookie.Value
			if value=="123"{
				c.Next()
				return
			}
		}
		c.Abort()
		return
	}
}