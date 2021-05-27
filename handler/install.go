package handler

import (
	"log"
	"net/http"
	"strings"

	"gocms/model"
	"gocms/pkg/config"

	"github.com/gin-gonic/gin"
)

// Install 安装配置
func Install(path string, srv *http.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := &model.Admin{
			Email:    strings.ToLower(c.PostForm("email")),
			Password: strings.TrimSpace(c.PostForm("password")),
			LastIP:   c.ClientIP(),
			Group:    model.Group{Name: "Administrator"},
		}

		err := c.ShouldBind(config.DB())
		if err != nil {
			c.Error(err)
			return
		}
		if err = model.Open(); err != nil {
			c.Error(err)
			return
		}
		if err = model.Install(u, path); err != nil {
			c.Error(err)
			return
		}

		c.Redirect(http.StatusFound, "/login")

		go func() {
			if err := srv.Shutdown(c); err != nil {
				log.Printf("Unable to shutdown the install server! Error: %v", err)
			}
		}()
	}
}
