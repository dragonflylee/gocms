package handler

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dragonflylee/gocms/model"
	"github.com/dragonflylee/gocms/util"
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
		Error(w, http.StatusNotFound, "页面错误 %v", err)
	} else if cookie, exist := session.Values["user"]; !exist {
		Error(w, http.StatusNotFound, "页面错误")
	} else if user, ok := cookie.(*model.Admin); !ok {
		Error(w, http.StatusNotFound, "页面错误")
	} else if user.Password = r.PostFormValue("password"); len(user.Password) < 8 {
		jFailed(w, http.StatusBadRequest, "密码不能少于8个字符")
	} else if err = user.UpdatePasswd(); err != nil {
		jFailed(w, http.StatusInternalServerError, err.Error())
	} else {
		session.Values["user"] = user
		session.Save(r, w)
		aLog(r, "修改管理员密码")
		jFailed(w, http.StatusOK, "修改密码成功")
	}
}

// Users 用户管理
func Users(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		Error(w, http.StatusBadRequest, err.Error())
		return
	}
	filter := func(db *gorm.DB) *gorm.DB {
		if email := strings.TrimSpace(r.Form.Get("email")); len(email) > 0 {
			db = db.Where("email = ?", strings.ToLower(email))
		}
		if group, err := strconv.ParseInt(r.Form.Get("group"), 10, 64); err == nil {
			db = db.Where("group_id = ?", group)
		}
		return db
	}
	data := make(map[string]interface{})
	if groups, err := model.GetGroups(); err == nil {
		data["group"] = groups
	}
	// 获取用户总数
	if nums, err := model.GetAdminNum(filter); err == nil && nums > 0 {
		p := util.NewPaginator(r, nums)
		if list, err := model.GetAdmins(func(db *gorm.DB) *gorm.DB {
			return db.Offset(p.Offset()).Limit(p.PerPageNums)
		}, filter); err == nil {
			data["list"] = list
		}
		data["page"] = p
	}
	rLayout(w, r, "users.tpl", data)
}

// UserAdd 用户添加
func UserAdd(w http.ResponseWriter, r *http.Request) {
	var (
		user model.Admin
		err  error
	)
	if err = r.ParseForm(); err != nil {
		jFailed(w, http.StatusBadRequest, err.Error())
		return
	}
	if user.Email = strings.ToLower(r.PostForm.Get("email")); len(user.Email) < 0 {
		jFailed(w, http.StatusBadRequest, "邮箱非法")
		return
	}
	if user.GroupID, err = strconv.ParseInt(r.PostForm.Get("group"), 10, 64); err != nil {
		jFailed(w, http.StatusBadRequest, "用户组非法")
		return
	}
	if err = user.Create(); err != nil {
		jFailed(w, http.StatusInternalServerError, err.Error())
		return
	}
	aLog(r, "添加管理员: %s", user.Email)
	jSuccess(w, nil)
}

// UserDelete 用户删除
func UserDelete(w http.ResponseWriter, r *http.Request) {
	var (
		vars = mux.Vars(r)
		user model.Admin
		err  error
	)
	if user.ID, err = strconv.ParseInt(vars["id"], 10, 64); err != nil {
		jFailed(w, http.StatusBadRequest, err.Error())
		return
	}
	if err = user.Delete(); err != nil {
		jFailed(w, http.StatusInternalServerError, err.Error())
		return
	}
	aLog(r, "删除管理员: %d", user.ID)
	jSuccess(w, nil)
}

// GroupEdit 角色管理
func GroupEdit(w http.ResponseWriter, r *http.Request) {
	var (
		vars  = mux.Vars(r)
		group model.Group
		err   error
	)
	if r.Method == http.MethodGet {
		if group.ID, err = strconv.ParseInt(vars["id"], 10, 64); err != nil {
			Error(w, http.StatusBadRequest, err.Error())
			return
		}
		data := make(map[string]interface{})
		if err = group.Select(); err == nil {
			data["group"] = &group
		}
		if nodes, err := model.GetNodeAllNodes(); err == nil {
			data["node"] = nodes
		}
		t.ExecuteTemplate(w, "group.tpl", data)
		return
	}
	if group.ID, err = strconv.ParseInt(vars["id"], 10, 64); err != nil {
		jFailed(w, http.StatusBadRequest, err.Error())
		return
	}
	if err = r.ParseForm(); err != nil {
		jFailed(w, http.StatusBadRequest, err.Error())
		return
	}
	if group.Name = strings.TrimSpace(r.PostForm.Get("name")); len(group.Name) <= 0 {
		jFailed(w, http.StatusBadRequest, "用户组不能为空")
		return
	}
	if nodes, exist := r.PostForm["node"]; exist {
		group.Nodes = make([]*model.Node, 0, len(nodes))
		for _, id := range nodes {
			var n model.Node
			if n.ID, err = strconv.ParseInt(id, 10, 64); err == nil {
				group.Nodes = append(group.Nodes, &n)
			}
		}
	}
	if err = group.Update(); err != nil {
		jFailed(w, http.StatusInternalServerError, err.Error())
		return
	}
	aLog(r, "修改角色: %s", group.Name)
	jSuccess(w, nil)
}

// GroupAdd 添加角色
func GroupAdd(w http.ResponseWriter, r *http.Request) {
	var (
		group model.Group
		err   error
	)
	if err = r.ParseForm(); err != nil {
		jFailed(w, http.StatusBadRequest, err.Error())
		return
	}
	if group.Name = strings.TrimSpace(r.PostForm.Get("name")); len(group.Name) <= 0 {
		jFailed(w, http.StatusBadRequest, "用户组不能为空")
		return
	}
	if err = group.Create(); err != nil {
		jFailed(w, http.StatusInternalServerError, err.Error())
		return
	}
	aLog(r, "新增角色: %s", group.Name)
	jSuccess(w, nil)
}

// Logs 操作日志
func Logs(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		Error(w, http.StatusBadRequest, err.Error())
		return
	}
	filter := func(db *gorm.DB) *gorm.DB {
		if id, err := strconv.ParseInt(r.Form.Get("id"), 10, 64); err == nil {
			db = db.Where("admin_id = ?", id)
		}
		if from, err := time.Parse(dateFormate, r.Form.Get("from")); err == nil {
			db = db.Where("created_at >= ?", from)
		}
		if to, err := time.Parse(dateFormate, r.Form.Get("to")); err == nil {
			db = db.Where("created_at < ?", to.AddDate(0, 0, 1))
		}
		return db
	}
	data := make(map[string]interface{})
	if _, exist := r.Form["export"]; exist {
		if list, err := model.GetLogs(filter); err == nil {
			data["日志列表"] = list
		}
		util.Excel(w, data, "操作日志 %s.xlsx", time.Now().Format(dateFormate))
		return
	}
	// 获取用户总数
	if nums, err := model.GetLogNum(filter); err == nil && nums > 0 {
		p := util.NewPaginator(r, nums)
		if list, err := model.GetLogs(func(db *gorm.DB) *gorm.DB {
			return db.Offset(p.Offset()).Limit(p.PerPageNums)
		}, filter); err == nil {
			data["list"] = list
		}
		data["page"] = p
	}
	rLayout(w, r, "logs.tpl", data)
}
