package handler

import (
	"net/http"

	"gocms/model"
	"gocms/pkg/auth"

	"github.com/gin-gonic/gin"
)

const (
	dateFormate = "2006-01-02"

	userKey = "User"
	menuKey = "Menu"
	nodeKey = "Node"
)

// Auth 认证
func Auth(c *gin.Context) {
	user := auth.SignedIn(c.Request, c.Writer)
	if user == nil {
		c.Redirect(http.StatusFound, "/login")
		c.Abort()
		return
	}
	// find node by route path
	node := model.GetNodeByPath(c.FullPath())
	if node == nil {
		node = new(model.Node)
	} else if !node.HasGroup(user.GroupID) {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	c.Set(userKey, user)
	c.Set(menuKey, model.GetNodes())
	c.Set(nodeKey, node)

	c.Next()

	if r := c.Request; r.Method == http.MethodPost {
		// 日志记录
		model.RecordLog(&model.AdminLog{
			AdminID: user.ID,
			Path:    r.RequestURI,
			UA:      r.UserAgent(),
			IP:      c.ClientIP(),
			Commit:  node.Name,
		})
	}
}
