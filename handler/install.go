package handler

import (
	"gocms/model"
	"gocms/util"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// Install 安装配置
func Install(path string, debug bool, s *mux.Router) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			t.ExecuteTemplate(w, "install.tpl", nil)
			return
		}
		if err := r.ParseForm(); err != nil {
			jFailed(w, http.StatusInternalServerError, err.Error())
			return
		}
		user := &model.Admin{
			Email:    strings.ToLower(strings.TrimSpace(r.PostForm.Get("email"))),
			Password: strings.TrimSpace(r.PostForm.Get("password")),
			Headpic:  "/static/img/avatar.png",
			Group:    model.Group{Name: "超级管理员"},
			LastIP:   r.RemoteAddr,
			Status:   true,
		}
		if !emailRegexp.MatchString(user.Email) {
			jFailed(w, http.StatusBadRequest, "邮箱格式非法")
			return
		}
		if !md5Regexp.MatchString(user.Password) {
			jFailed(w, http.StatusBadRequest, "密码不正确")
			return
		}
		if err := util.ParseForm(r.PostForm, &model.Config.DB); err != nil {
			jFailed(w, http.StatusBadRequest, err.Error())
			return
		}
		if err := model.Open(debug); err != nil {
			jFailed(w, http.StatusInternalServerError, err.Error())
			return
		}
		if err := model.Install(user, path); err != nil {
			jFailed(w, http.StatusInternalServerError, err.Error())
			return
		}
		jSuccess(w, "/login")
	})
}
