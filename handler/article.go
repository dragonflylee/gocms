package handler

import (
	"fmt"
	"gocms/model"
	"gocms/util"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// Articles 文章列表
func Articles(w http.ResponseWriter, r *http.Request) {
	filter := func(db *gorm.DB) *gorm.DB {
		if from, err := time.ParseInLocation(dateFormate, r.Form.Get("from"), time.Local); err == nil {
			db = db.Where("created_at >= ?", from)
		}
		if to, err := time.ParseInLocation(dateFormate, r.Form.Get("to"), time.Local); err == nil {
			db = db.Where("created_at <= ?", to)
		}
		return db
	}
	data := make(map[string]interface{})
	if nums, err := model.GetArticleNum(filter); err == nil && nums > 0 {
		p := util.NewPaginator(r, nums)
		if list, err := model.GetArticles(func(db *gorm.DB) *gorm.DB {
			return db.Offset(p.Offset()).Limit(p.PerPageNums)
		}, filter); err == nil {
			data["List"] = list
		}
		data["Page"] = p
	}
	rLayout(w, r, "articles.tpl", data)
}

// GetArticle 文章详情
func GetArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var v model.Article
	if err = v.Query(id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	rLayout(w, r, "article_detail.tpl", v)
}

// EditArticle 编辑文章
func EditArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	v := &model.Article{ID: id, Title: "新建文章"}
	// 如果是get请求，则跳转到编辑页
	if r.Method == http.MethodGet {
		if err = v.Query("id = ?", id); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		rLayout(w, r, "article_edit.tpl", v)
		return
	}
	if err = r.ParseForm(); err != nil {
		jFailed(w, http.StatusBadRequest, err.Error())
		return
	}
	v.Title = r.PostForm.Get("title")
	v.Content = r.PostForm.Get("content")

	if err = v.Update(); err != nil {
		jFailed(w, http.StatusBadRequest, err.Error())
		return
	}
	jSuccess(w, fmt.Sprintf("/article/%d", v.ID))
}
