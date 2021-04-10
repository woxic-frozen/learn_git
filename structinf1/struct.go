package strcutinf
//用户结构体
type User struct {
	Id string `form:"id"`
	Password string `form:"password"`
}
//文章结构体
type ArticleInfo struct{
	Title string `form:"title" binding:"required"`//文章标题
	Context string `form:"context" binding:"required"`//文章正文
	Id string `form:"id"`//作者id
	likes int//点赞数
}
type Message struct{
	Aid int  `form:"aid"`//文章表自增项为每篇文章唯一
	Mid int `form:"mid"`//评论自增项为每个评论唯一
	Message string `form:"message"`//评论内容
	Id string `form:"id"`//
}