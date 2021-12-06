package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 数据库连接 start
var db *gorm.DB

type (
	Tag struct {
		ID   int `gorm: "primaryKey"`
		Name string
	}
	User struct {
		gorm.Model
		Name     string  `json: "name" binding: "required" gorm: "uniqueIndex"`
		Phone    string  `json: "phone" binding: "requied" gorm: "uniqueIndex"`
		Password string  `json: "password" binding: "required"`
		Avatar   string  `json: "avatar"`
		Gender   int8    `json: "gender"`
		Birthday int64   `json: "birthday"`
		Email    *string `json: "email"`
	}
	Article struct {
		gorm.Model
		Title    string `json: "title"`
		Content  string `gorm: "size: 1024"`
		AuthorID int
		Author   User `gorm: "foreignKey:AuthorID"`
		Tag      Tag  `gorm: "embedded;embeddedPrefix:tag_` // ?
	}
)

func init() {
	var err error
	dsn := "root:123456@tcp(127.0.0.1:3306)/demo?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic("failed to connect database:" + err.Error())
	}
	db.AutoMigrate(&User{}, &Article{}, &Tag{})
}

// 数据库连接 end

// 跨域中间件
func CrosHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Origin", "*") // 设置允许访问所有域
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
		context.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma,token,openid,opentoken")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
		context.Header("Access-Control-Max-Age", "172800")
		context.Header("Access-Control-Allow-Credentials", "true")
		context.Set("content-type", "application/json") // 设置返回格式是json
		//处理请求
		context.Next()
	}
}

type Login struct {
	Username string `form: "username" json: "username" binding: "required"`
	Password string `form: "password" json: "password" binding: "required"`
}

func main() {
	r := gin.Default()
	r.Static("/static", "./static") // 设置静态资源目录
	r.Use(CrosHandler())            // 允许跨域中间件
	r.MaxMultipartMemory = 20       // 文件上传最大尺寸
	r1 := r.Group("/day1")          // 接口分组
	r2 := r.Group(("/day2"))        // 接口分组
	r3 := r.Group(("/day3"))        // 接口分组

	/* =========================== day1 =========================== */
	// 路由
	r1.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// 路由uri参数
	r1.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		// c.String(200, name)
		c.JSON(200, gin.H{
			"name":   name,
			"action": action,
		})
	})
	// 路由query参数
	r1.GET("/student", func(c *gin.Context) {
		id := c.Query("id")
		c.String(200, id)
	})
	// 路由body参数
	r1.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.JSON(200, gin.H{
			"code":    0,
			"message": "success",
			"data": gin.H{
				"username": username,
				"password": password,
			},
		})
	})
	// 单文件上传
	r1.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file)
		c.SaveUploadedFile(file, "files/"+file.Filename)
		c.JSON(200, gin.H{
			"fileName": file.Filename,
		})
	})
	// 多文件上传
	r1.POST("/uploads", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["files[]"]
		for _, file := range files {
			log.Println(file.Filename)
			c.SaveUploadedFile(file, "files/"+file.Filename)
		}
		c.String(200, "save success")
	})

	/* =========================== day2 =========================== */
	// 参数校验
	r2.POST("/loginJson", func(c *gin.Context) {
		var reqJson Login
		if err := c.ShouldBindJSON(&reqJson); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": err.Error(),
			})
			return
		}
		if reqJson.Username == "admin" && reqJson.Password == "123" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    403,
				"message": "账号或者密码错误",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "success",
			"data": gin.H{
				"username": reqJson.Username,
				"password": reqJson.Password,
			},
		})
	})

	r.LoadHTMLGlob("assets/templates/*") // 设置网页模板目录

	// 网页模板与数据传输
	r2.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "hello ksir",
		})
	})
	// r.LoadHTMLFiles("assets/html/index.html", "assets/html/home.html") // 设置网页目录
	// r2.GET("/index.html", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.html", gin.H{
	// 		"title": "hello go",
	// 	})
	// })
	// r2.GET("/home.html", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "home.html", gin.H{
	// 		"title": "hello go",
	// 	})
	// })

	/* =========================== day3 =========================== */
	r3.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.tmpl", gin.H{})
	})
	r3.POST("/register", func(c *gin.Context) {
		var reqData User
		if err := c.ShouldBindJSON(&reqData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": err.Error(),
			})
			return
		} else if len(reqData.Name) < 1 || len(reqData.Phone) < 11 || len(reqData.Password) < 6 {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "请检查数据完整性",
			})
			return
		}
		result := db.Create(&reqData)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "数据库存储异常",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "注册成功",
			"data": gin.H{
				"name":     reqData.Name,
				"phone":    reqData.Phone,
				"password": reqData.Password,
			},
		})
	})
	r3.GET("/delete", func(c *gin.Context) {
		userId := c.Query("id")
		result := db.Delete(&User{}, userId)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "数据库删除异常",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "删除成功",
		})
	})
	r3.GET("/updateUser", func(c *gin.Context) {
		userId := c.Query("id")
		var user User
		db.First(&user, userId)
		log.Println(user)
		c.HTML(http.StatusOK, "updateUser.tmpl", user)
	})
	r3.POST("/updateUser", func(c *gin.Context) {
		var reqData User
		if err := c.ShouldBindJSON(&reqData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": err.Error(),
			})
			return
		} else if len(reqData.Name) < 1 || len(reqData.Phone) < 11 || len(reqData.Password) < 6 {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "请检查数据完整性",
			})
			return
		}
		result := db.Save(&reqData)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "数据库存储异常",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "更新成功",
		})
	})
	r.Run()
}
