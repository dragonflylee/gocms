package handler

import (
	"net/http"
	"path/filepath"

	"github.com/dragonflylee/gocms/model"
	"github.com/gorilla/mux"
)

// Install 安装配置
func Install(path string, t *mux.Router) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			render(w, r, "install.tpl", nil)
			return
		}
		if err := r.ParseForm(); err != nil {
			rsp(w, http.StatusInternalServerError, err.Error(), nil)
			return
		}
		c := &model.Config{
			Host: r.FormValue("host"),
			User: r.FormValue("user"),
			Pass: r.FormValue("pass"),
			Name: r.FormValue("name"),
		}
		if err := model.Open(c); err != nil {
			rsp(w, http.StatusInternalServerError, err.Error(), nil)
			return
		}
		if err := c.Save(filepath.Join(path, "config.json")); err != nil {
			rsp(w, http.StatusInternalServerError, err.Error(), nil)
			return
		}
		u := &model.Admin{
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
			Email:    r.FormValue("email"),
			Group:    model.Group{Name: "超级管理员"},
		}
		if err := u.Group.Create(); err != nil {
			rsp(w, http.StatusInternalServerError, err.Error(), nil)
			return
		}
		if err := u.Create(); err != nil {
			rsp(w, http.StatusInternalServerError, err.Error(), nil)
			return
		}
		Route(t)
		rsp(w, http.StatusOK, "安装成功", "/login")
	})
}
