package handler

import (
	"net/http"
	"strings"

	"github.com/dragonflylee/gocms/model"
	"github.com/jinzhu/gorm"
)

// Home 首页
func Home(w http.ResponseWriter, r *http.Request) {
	rLayout(w, r, "index.tpl", nil)
}

// Profile 个人中心
func Profile(w http.ResponseWriter, r *http.Request) {
	rLayout(w, r, "profile.tpl", nil)
}

// Users 用户管理
func Users(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	filter := func(db *gorm.DB) *gorm.DB {
		if email := strings.TrimSpace(r.Form.Get("email")); email != "" {
			db = db.Where("admins.email = ?", strings.ToLower(email))
		}
		return db
	}
	// 获取用户总数
	nums, err := model.GetAdminNum(filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if nums <= 0 {
		rLayout(w, r, "users.tpl", emptyData)
		return
	}
	p := NewPaginator(r, nums)
	users, err := model.GetAdmins(p.PerPageNums, p.Offset(), filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	groups, err := model.GetGroups()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for _, u := range users {
		u.Group.Name = groups[u.GroupID]
	}
	rLayout(w, r, "users.tpl", map[string]interface{}{
		"list": users, "groups": groups, "page": p})
}

// Logs 操作日志
func Logs(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	filter := func(db *gorm.DB) *gorm.DB {
		if email := strings.TrimSpace(r.Form.Get("email")); email != "" {
			db = db.Where("admins.email = ?", strings.ToLower(email))
		}
		return db
	}
	// 获取用户总数
	nums, err := model.GetLogNum(filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if nums <= 0 {
		rLayout(w, r, "logs.tpl", emptyData)
		return
	}
	p := NewPaginator(r, nums)
	logs, err := model.GetLogs(p.PerPageNums, p.Offset(), filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rLayout(w, r, "logs.tpl", map[string]interface{}{
		"list": logs, "page": p})
}
