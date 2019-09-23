package main

import (
	//_ "net/http/pprof"
	"github.com/DeanThompson/ginpprof"

	"blog/controller"
	"blog/dal/db"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//在连接的dsn中，添加parseTime=true 和loc=Local，此处的local可以换为具体的时区(Asia/Shanghai)
	dns := "root:123456@tcp(localhost:3306)/blogger?parseTime=true&loc=Local"
	err := db.Init(dns)
	if err != nil {
		panic(err)
	}

	ginpprof.Wrapper(router)
	router.Static("/static/", "./static")
	router.LoadHTMLGlob("views/*")

	router.GET("/", controller.IndexHandle)
	//发布文章页面
	router.GET("/article/new/", controller.NewArticle)
	//文章提交接口
	router.POST("/article/submit/", controller.ArticleSubmit)
	//文章详情页
	router.GET("/article/detail/", controller.ArticleDetail)

	//文件上传接口
	router.POST("/upload/file/", controller.UploadFile)

	//留言页面
	router.GET("/leave/new/", controller.LeaveNew)
	//关于我页面
	router.GET("/about/me/", controller.AboutMe)

	//文章评论相关
	router.POST("/comment/submit/", controller.CommentSubmit)

	//留言相关
	router.POST("/leave/submit/", controller.LeaveSubmit)
	//分类下面的文章列表
	router.GET("/category/", controller.CategoryList)
	router.Run(":8080")
}
