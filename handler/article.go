package handler

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"gocms/model"
	"gocms/pkg/errors"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/russross/blackfriday/v2"
)

// Articles 文章列表
func Articles(c *gin.Context) {
	req := new(model.DateRangeOpts)

	if err := c.ShouldBindUri(req); err != nil {
		return
	}

	if nums, err := model.GetArticleNum(req); err == nil && nums > 0 {
		p := model.NewPaginator(c.Request.URL, nums)
		if list, err := model.GetArticles(req, p); err == nil {
			c.Set("List", list)
		}
		c.Set("Page", p)
	}
	c.HTML(http.StatusOK, "articles.html", c.Keys)
}

// GetArticle 文章详情
func GetArticle(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}
	var v model.Article
	if err = v.Query(id); err != nil {
		c.Error(err)
		return
	}
	md := []byte(strings.ReplaceAll(v.Content, "\r\n", "\n"))
	c.Set("Data", template.HTML(blackfriday.Run(md)))
	c.HTML(http.StatusOK, "article_detail.html", c.Keys)
}

// EditArticle 编辑文章
func EditArticle(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.Error(err)
		return
	}
	v := &model.Article{ID: id, Title: "新建文章"}
	// 如果是get请求，则跳转到编辑页
	if c.Request.Method == http.MethodGet {
		if v.ID > 0 {
			if err = v.Query(); err != nil {
				c.Error(err)
				return
			}
		}
		c.Set("Data", v)
		c.HTML(http.StatusOK, "article_edit.html", c.Keys)
		return
	}

	if err = c.MustBindWith(v, binding.FormPost); err != nil {
		return
	}

	v.AdminID = c.MustGet(userKey).(*model.Admin).ID

	if err = v.Update(); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, errors.OK())
}
