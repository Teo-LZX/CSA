package main
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)
var db *gorm.DB

type Comment struct{     //一个评论结构体
	ID string `gorm: "primary key"`          //评论id
	Name string `gorm: "not null"`   //用户名
	Content string `gorm: "not null"` //评论内容
	Zan int `gorm: "default:0"`        //此评论获得的赞
	Pre_ID string      //此评论所评论的对象，若没有则为空
}
type UserInfo struct {
	UserName string
	Password string
}

func init_mysql()(err error){  //初始化数据库
	dsn := "Teo:635695606@(localhost)/message_board?charset=utf8mb4&parseTime=True&loc=Local"
	db,err = gorm.Open("mysql", dsn)
	if err!=nil{
		fmt.Println(err)
		return
	}
	return db.DB().Ping()
}

func SignUp(c *gin.Context,){   //注册
	UserName := c.Param("name")
	Password := c.Param("psd")
	user := UserInfo{UserName, Password}
	err := db.Create(&user).Error
	if err!=nil{
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}else{
		c.HTML(http.StatusOK, "view.html", gin.H{
			"template": "注册成功，请保管好密码!",
		})
	}
}

func SignIn(c *gin.Context){  //登录
	var user UserInfo    //用来存放用户输入的用户名和密码
	var user2 UserInfo  //存放从数据库中查找到的用户名和密码
	user.UserName = c.Param("username")    //解析参数，存入user中
	user.Password = c.Param("password")
	err := db.Where("user_name=?", user.UserName).First(&user2).Error  //将数据库中所有与输入的uername比对，如果没有返回错误
	if err!=nil || user2.Password!= user.Password{   //用户名相同，密码错误或者用户名不存在
		c.HTML(http.StatusOK, "view.html", gin.H{
			"template": "用户名或密码错误",
		})
	}else{   //否侧返回登录成功信息
		c.HTML(http.StatusOK, "view.html", gin.H{
			"template": "登录成功!",
		})
	}
}

func AddComment(c * gin.Context){   //发布评论
	//前端发请求后到这里，从请求把数据拿出来，存入数据库，返回响应
	var comment Comment
	//处理参数信息
	comment.ID = c.Query("id")
	comment.Name = c.Query("name")
	comment.Content = c.Query("content")
	comment.Pre_ID = c.Query("pre")
	fmt.Println(comment)
	//c.BindJSON(&comment)
	err := db.Create(&comment).Error //将信息存入数据库
	if err != nil{
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}else{
		c.JSON(http.StatusOK, comment)
		fmt.Println("success")
	}
}

func View(c *gin.Context){  //查看所有评论
	var comments []Comment
	err := db.Find(&comments).Error
	if err!=nil{
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}else{
		c.JSON(http.StatusOK, comments)
		c.HTML(http.StatusOK, "view.html", nil)
	}
}

func Query(c *gin.Context){  //查询评论
	var comment Comment
	id := c.Param("id")
	err := db.Where("id=?", id).First(&comment).Error
	if err!=nil{
		c.JSON(http.StatusOK, gin.H{
			"error":err.Error(),
		})
	}else{
		c.JSON(http.StatusOK, comment)
		c.HTML(http.StatusOK, "view.html", nil)
	}
}

func Change(c *gin.Context){  //修改评论
	//-------------------解析参数--------------------------
	id := c.Param("id")
	content := c.Param("content")
	//---------------------修改数据库内容--------------------------
	var comment Comment
	comment.Content = content
	err := db.Model(&comment).Where("id=?", id).Update("content", content).Error
	if err!=nil{
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, comment)
}

func Delete(c *gin.Context){  //删除评论
	id := c.Param("id")
	db.Where("id=?", id).Delete(Comment{})
	c.HTML(http.StatusOK, "view.html", gin.H{
		"template": "删除成功!",
})
}

func Zan(c *gin.Context){
	var comment Comment
	id := c.Param("id")
	err := db.Model(&comment).Where("id=?", id).First(&comment).Error
	comment.Zan += 1
	db.Save(&comment)
	if err!=nil{
		c.JSON(http.StatusOK, gin.H{
			"error":err.Error(),
		})
	}else{
		c.JSON(http.StatusOK, comment)
		c.HTML(http.StatusOK, "view.html", gin.H{
			"templates": "点赞成功",
		})
	}
}

func main() {
//-------------------数据库--------------------------
	err := init_mysql() //创建并连接数据库
	if err != nil {
		return
	}
	defer db.Close()  //程序结束关闭数据库
	db.AutoMigrate(&Comment{}) //模型绑定
	db.AutoMigrate(&UserInfo{})
//---------------------启动gin------------------------
	r := gin.Default()    //开一个路由
	r.LoadHTMLGlob("./templates/*")//加载模板文件
	//r.Static("/static", "static")
	//r.GET("/", get_index)
//---------------------------------------------
	HMP := r.Group("sb")
	{
		HMP.POST("/signup/:name/:psd", SignUp) //注册

		HMP.POST("/signin/:username/:password", SignIn)  //登录

		HMP.POST("/", AddComment)  /*发布评论*/

		HMP.GET("/query", View)   //查询看所有评论

		HMP.GET("/query/:id", Query)    //查询评论

		HMP.GET("/zan/:id", Zan)    //点赞

		HMP.PUT("/change/:id/:content",  Change)    //修改评论

		HMP.DELETE("/delete/:id", Delete)    //删除评论
	}
	r.Run(":9000")
}
