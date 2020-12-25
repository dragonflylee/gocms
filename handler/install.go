package handler

import (
	"gocms/model"
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
			jFailed(w, http.StatusBadRequest, err.Error())
			return
		}
		u := &model.Admin{
			Email:    strings.TrimSpace(r.PostForm.Get("email")),
			Password: strings.TrimSpace(r.PostForm.Get("password")),
			Headpic:  "/static/img/avatar.png", LastIP: r.RemoteAddr,
			Group: model.Group{Name: "超级管理员"},
		}
		if u.Email = strings.ToLower(u.Email); !emailRegexp.MatchString(u.Email) {
			jFailed(w, http.StatusBadRequest, "邮箱格式非法")
			return
		}
		if err := decoder.Decode(&model.Config.DB, r.PostForm); err != nil {
			jFailed(w, http.StatusBadRequest, err.Error())
			return
		}
		if err := model.Open(debug); err != nil {
			jFailed(w, http.StatusInternalServerError, err.Error())
			return
		}
		if err := model.Install(u, path); err != nil {
			jFailed(w, http.StatusInternalServerError, err.Error())
			return
		}
		jSuccess(w, "/login")
	})
}
