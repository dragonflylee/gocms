package handler

import (
	"net/http"
	"strconv"
	"strings"

	"gocms/model"
	"gocms/pkg/config"
	"gocms/pkg/errors"
	"gocms/pkg/util"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// Profile 个人中心
func Profile(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		c.Set("Captcha", config.Captcha())
		c.HTML(http.StatusOK, "profile.html", c.Keys)
		return
	}

	u := c.MustGet(userKey).(*model.Admin)

	switch c.Query("action") {
	case "password":
		if u.Password = c.PostForm("password"); len(u.Password) < 8 {
			c.Error(errors.New("PasswordTooShort")) // 密码不能少于8个字符
			return
		}
		u.Salt = util.RandString(5)
		u.Password = util.MD5(u.Password, util.MD5(u.Salt))
		u.Flags = (u.Flags & ^model.FlagResetPassNext) | model.FlagPassNeverExpire

		if err := u.Update("password", "salt", "flags"); err != nil {
			c.Error(err)
			return
		}

	case "phone":

		code := strings.TrimSpace(c.PostForm("code"))
		if len(code) < 1 {
			c.Error(errors.ErrCapcha)
			return
		}

	default:
		c.Error(errors.ErrRequest)
		return
	}

	c.JSON(http.StatusOK, errors.OK())
}

// Users 用户管理
func Users(c *gin.Context) {
	if groups, err := model.GetGroups(); err == nil {
		c.Set("Group", groups)
	}
	// 获取用户总数
	if nums, err := model.GetAdminNum(); err == nil && nums > 0 {
		p := model.NewPaginator(c.Request.URL, nums)
		if list, err := model.GetAdmins(p); err == nil {
			c.Set("List", list)
		}
		c.Set("Page", p)
	}
	c.HTML(http.StatusOK, "users.html", c.Keys)
}

// UserAdd 用户添加
func UserAdd(c *gin.Context) {
	var req struct {
		Email string `form:"email" binding:"email,required"`
		Group int64  `form:"group" binding:"required"`
	}

	err := c.MustBindWith(&req, binding.FormPost)
	if err != nil {
		return
	}

	u := c.MustGet(userKey).(*model.Admin)
	if req.Group == u.GroupID {
		c.Error(errors.ErrForbbiden)
		return
	}

	u = &model.Admin{Email: req.Email, GroupID: req.Group}
	if err := model.NewAdmin(u); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, errors.OK())
}

// UserDelete 用户删除
func UserDelete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}
	if err = model.DeleteAdmin(&model.Admin{ID: id}); err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}
	c.JSON(http.StatusOK, errors.OK())
}

// GroupEdit 角色管理
func GroupEdit(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	if c.Request.Method == http.MethodGet {
		group, err := model.GetGroup(id)
		if err != nil {
			c.Error(err)
			return
		}
		group.Nodes = model.GetNodes()
		c.HTML(http.StatusOK, "group.html", group)
		return
	}

	u := c.MustGet(userKey).(*model.Admin)
	if id == u.GroupID {
		c.Error(errors.ErrForbbiden)
		return
	}

	group := new(model.Group)
	err = c.MustBindWith(&group, binding.FormPost)
	if err != nil {
		return
	}

	if nodes, exist := c.GetQueryArray("node[]"); exist {
		group.Nodes = make([]*model.Node, 0, len(nodes))
		for _, id := range nodes {
			var n model.Node
			if n.ID, err = strconv.Atoi(id); err == nil {
				group.Nodes = append(group.Nodes, &n)
			}
		}
	}
	if err = group.Update(); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, errors.OK())
}

// GroupAdd 添加角色
func GroupAdd(c *gin.Context) {
	var group model.Group

	err := c.MustBindWith(&group, binding.FormPost)
	if err != nil {
		return
	}
	if err = group.Create(); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, errors.OK())
}

// Logs 操作日志
func Logs(c *gin.Context) {
	req := new(model.DateRangeOpts)

	if err := c.ShouldBindUri(req); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	// 获取用户总数
	if nums, err := model.GetLogNum(req); err == nil && nums > 0 {
		p := model.NewPaginator(c.Request.URL, nums)
		if list, err := model.GetLogs(req, p); err == nil {
			c.Set("List", list)
		}
		c.Set("Page", p)
	}
	c.HTML(http.StatusOK, "logs.html", c.Keys)
}
