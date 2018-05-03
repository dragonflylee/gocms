package handler

import (
	"bytes"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/dragonflylee/gocms/model"
	"github.com/gorilla/mux"
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

// Password 密码修改
func Password(w http.ResponseWriter, r *http.Request) {
	if session, err := store.Get(r, sessName); err != nil {
		http.NotFound(w, r)
	} else if cookie, exist := session.Values["user"]; !exist {
		http.NotFound(w, r)
	} else if user, ok := cookie.(*model.Admin); !ok {
		http.NotFound(w, r)
	} else if user.Password = r.PostFormValue("passwd"); len(user.Password) < 8 {
		jRsp(w, http.StatusBadRequest, "密码不能少于8个字符", nil)
	} else if err = user.UpdatePasswd(); err != nil {
		jRsp(w, http.StatusInternalServerError, err.Error(), nil)
	} else {
		session.Values["user"] = user
		session.Save(r, w)
		aLog(r, "修改管理员密码")
		jRsp(w, http.StatusOK, "修改成功", nil)
	}
}

// Users 用户管理
func Users(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	filter := func(db *gorm.DB) *gorm.DB {
		if email := strings.TrimSpace(r.Form.Get("email")); email != "" {
			db = db.Where("email = ?", strings.ToLower(email))
		}
		if group, err := strconv.ParseInt(r.Form.Get("group"), 10, 64); err == nil {
			db = db.Where("group_id = ?", group)
		}
		return db
	}
	groups, err := model.GetGroups()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// 获取用户总数
	nums, err := model.GetAdminNum(filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if nums <= 0 {
		rLayout(w, r, "users.tpl", map[string]interface{}{"list": nil, "group": groups, "page": nil})
		return
	}
	p := NewPaginator(r, nums)
	users, err := model.GetAdmins(p.PerPageNums, p.Offset(), filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for _, u := range users {
		u.Group.Name = groups[u.GroupID]
	}
	rLayout(w, r, "users.tpl", map[string]interface{}{
		"list": users, "group": groups, "page": p})
}

// UserAdd 用户添加
func UserAdd(w http.ResponseWriter, r *http.Request) {
	var (
		user model.Admin
		body bytes.Buffer
		err  error
	)
	if err = r.ParseForm(); err != nil {
		jRsp(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	if user.Email = strings.ToLower(r.PostForm.Get("email")); len(user.Email) < 0 {
		jRsp(w, http.StatusBadRequest, "邮箱非法", nil)
		return
	}
	if user.GroupID, err = strconv.ParseInt(r.PostForm.Get("group"), 10, 64); err != nil {
		jRsp(w, http.StatusBadRequest, "用户组非法", nil)
		return
	}
	user.Password = fmt.Sprint(rand.Intn(8999999) + 1000000)
	if err = t.ExecuteTemplate(&body, "email.tpl", &user); err != nil {
		jRsp(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	if err = user.Create(); err != nil {
		jRsp(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	aLog(r, fmt.Sprintf("添加管理员 (%s)", user.Email))
	jRsp(w, http.StatusOK, "添加成功", nil)
}

// GroupEdit 角色管理
func GroupEdit(w http.ResponseWriter, r *http.Request) {
	var (
		vars    = mux.Vars(r)
		id, err = strconv.ParseInt(vars["id"], 10, 64)
	)
	if r.Method == http.MethodGet {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else if nodes, err := model.GetNodeAllNodes(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			t.ExecuteTemplate(w, "group.tpl", map[string]interface{}{
				"id": id, "node": nodes})
		}
		return
	}
	jRsp(w, http.StatusBadRequest, "无权操作", nil)
}

// GroupAdd 添加角色
func GroupAdd(w http.ResponseWriter, r *http.Request) {
	jRsp(w, http.StatusBadRequest, "无权操作", nil)
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
		rLayout(w, r, "logs.tpl", map[string]interface{}{"list": nil, "page": nil})
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
