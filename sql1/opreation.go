package sql

import (
	"database/sql"
	"ginblog/structinf1"
	_ "github.com/go-sql-driver/mysql"
)
var db *sql.DB
//连接操作
func OpenMysql()*sql.DB{
	db,_=sql.Open("mysql","root:@tcp(localhost:3306)/blog?charset=utf8")
	return db
	}
/*var DB *sql.DB
func init(){
	DB,_=sql.Open("mysql","root:191513@tcp(127.0.0.1:3306)/blog?charset=utf8")
	DB.SetMaxOpenConns(1000)
	err:=DB.Ping()
	if err!=nil{
		fmt.Println("fail to connect to db")
	}
}*/

//func OpenMysql()*sql.DB{
	//DB,_=sql.Open("mysql","root:191513@tcp(127.0.0.1:3306)/blog?charset=utf8")
	//DB.SetMaxOpenConns(1000)
	//err:=DB.Ping()
	//if err!=nil{
	//	fmt.Println("fail to connect to db")
	//}
	//return  DB
//}
///var db *sql.DB

//登录的sql操作
func Find(id1 string,password1 string)(decide bool){
	db,_:=sql.Open("mysql","root:@tcp(127.0.0.1:3306)/blog?charset=utf8")
	stmt,err:=OpenMysql().Query("SELECT id,password FROM usrtable")//select
	if err!=nil{
		panic(err)
	}
	defer db.Close()
	for stmt.Next(){
		var id string
		var password string
		err:=stmt.Scan(&id,&password)
		if err!=nil{
			panic(err)
		}
		if id==id1&&password==password1&&id!="" {
			/*stmt1,_:=db.Prepare("UPDATE ginexmp SET status=?WHERE id=?")
			stmt1.Exec(1,id1)*/
			return true
		}
	}
	return false
}

//注册的sql操作
func Register(id1 string,password string)bool{
	//db,_:=sql.Open("mysql","root:@tcp(127.0.0.1:3306)/blog?charset=utf8")
	stmt,_:=OpenMysql().Query("SELECT id FROM usrtable ")
	defer stmt.Close()
	//defer db.Close()
	for stmt.Next(){
		var id string
		stmt.Scan(&id)
		if id==id1{
			return false
		}
	}
	stmt1,err:=OpenMysql().Prepare("INSERT INTO usrtable (id,password)VALUES (?,?)")
	if err!=nil{
		panic(err)
	}
	defer stmt1.Close()
	stmt1.Exec(id1,password)
	return true
}
func UpArtile(id string,aid int,message string){
	//db,_:=sql.Open("mysql","root:@tcp(127.0.0.1:3306)/blog?charset=utf8")
	stmt,err:=OpenMysql().Prepare("INSERT INTO mtable (id,aid,message)VALUES(?,?,?)")
	defer db.Close()
	if err!=nil{
		panic(err)
	}
	stmt.Exec(id,aid,message)
}
func queryone(id string){
	//stmt,err:=db.Prepare("SELECT *FROM Atable")
	//if err!=nil{
		//panic(err)
	//}

}
//根据文章aid查找文章内容和评论
//输入aid int返回文章的结构体和评论结构体的切片
func QueryArticle(aid int)(strcutinf.ArticleInfo,[] strcutinf.Message){
	stmt,err:=OpenMysql().Prepare("SELECT id,title,context FROM atable WHERE aid=?")
	defer stmt.Close()
	if err!=nil{
		panic(err)
	}
	row:=stmt.QueryRow(aid)
	article:=strcutinf.ArticleInfo{}
	row.Scan(&article.Id,&article.Title,&article.Context)
	stmt1,err1:=db.Prepare("SELECT aid,message,id FROM mtable WHERE aid=?")
	defer stmt1.Close()
	if err1!=nil{
		panic(err)
	}
	rows,_:=stmt1.Query(aid)
	 message1:=make([]strcutinf.Message,0)
	for rows.Next(){
		var message strcutinf.Message
		err:=rows.Scan(&message.Aid,&message.Message,&message.Id)
		if err!=nil{
			panic(err)
		}
		message1=append(message1,message)
	}
	return article,message1
}
//发布文章
func Luancharticle(info strcutinf.ArticleInfo){
	db,_:=sql.Open("mysql","root:@tcp(127.0.0.1:3306)/blog?charset=utf8")
	stmt,err:=db.Prepare("INSERT INTO atable (title,context,id)VALUES (?,?,?)")
	if err!=nil{
		panic(err)
	}
	defer stmt.Close()
	defer db.Close()
	stmt.Exec(info.Title,info.Context,info.Id)
}
//发表评论
func Luanchmessge(message1 strcutinf.Message){
	db,_:=sql.Open("mysql","root:@tcp(127.0.0.1:3306)/blog?charset=utf8")
	stmt,err:=db.Prepare("INSERT INTO mtable (message,id,aid) VALUES (?,?,?)")
	defer stmt.Close()
	defer db.Close()
	if err!=nil{
		panic(err)
	}
	stmt.Exec(message1.Message,message1.Id,message1.Aid)
}
//点赞sql操作
func Likes(aid int){
	stmt,err:=OpenMysql().Prepare("UPDATE atable SET likes=likes+1 WHERE aid=?")
	defer stmt.Close()
	if err==nil{
		stmt.Exec(aid)
	}else{
		panic(err)
	}
}