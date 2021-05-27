package main

import (
	"embed"
	"flag"
	"log"
	"net/http"

	"gocms/handler"
	"gocms/model"
	"gocms/pkg/auth"
	"gocms/pkg/config"
	"gocms/pkg/i18n"
	"gocms/pkg/route"

	"github.com/gin-gonic/gin"
)

//go:embed assets/* views/admin/*.html i18n/*.ini
var content embed.FS

func main() {
	var p string
	flag.StringVar(&p, "c", "config.yml", "config file path")
	flag.Parse()

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	err := i18n.Load(content, "i18n/*.ini")
	if err != nil {
		log.Fatalf("local i18n: %v", err)
	}
	// 加载配置文件
	if err = config.Load(p); err != nil {
		r := route.New(content)
		r.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "install.html", c.Keys)
		})
		srv := &http.Server{
			Addr: config.HTTP().Addr, Handler: r,
		}
		r.POST("/install", handler.Install(p, srv))
		done := route.Listen(srv)
		select {
		case <-done:
			return
		default:
		}
	} else if err = model.Open(); err != nil {
		log.Panicf("open db failed: %v", err)
	}

	auth.Init()

	r := route.New(content)
	// 登录相关
	r.POST("/login", route.Limit(2), handler.Login)
	r.GET("/bingpic", route.Limit(2), handler.BingPic)
	r.GET("/avatar", handler.Avatar)

	r.GET("/login", handler.Login)
	r.GET("/logout", handler.Logout)
	// 后台主页
	s := r.Group("/admin", handler.Auth)
	// 系统管理
	s.GET("/users", handler.Users)
	s.POST("/user/add", handler.UserAdd)
	s.POST("/user/del/:id", handler.UserDelete)
	s.GET("/group/:id", handler.GroupEdit)
	s.POST("/group/add", handler.GroupAdd)
	s.GET("/logs", handler.Logs)
	// 文章管理
	s.GET("/articles", handler.Articles)
	s.GET("/article/edit/:id", handler.EditArticle)
	s.POST("/article/edit/:id", handler.EditArticle)
	// 个人中心
	s.GET("/profile", handler.Profile)
	s.POST("/profile", handler.Profile)
	// 文件上传
	s.GET("/upload", handler.Upload)
	s.GET("/", handler.Dashboard)

	<-route.Listen(&http.Server{
		Addr: config.HTTP().Addr, Handler: r,
	})
}
