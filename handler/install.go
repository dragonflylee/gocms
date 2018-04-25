package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/dragonflylee/gocms/model"
	"github.com/gorilla/mux"
)

// Install 安装配置
func Install(path string, route *mux.Router) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			t.ExecuteTemplate(w, "install.tpl", nil)
			return
		}
		err := r.ParseForm()
		if err != nil {
			jRsp(w, http.StatusInternalServerError, err.Error(), nil)
			return
		}
		conf := &model.Config{
			Type: strings.ToLower(r.PostForm.Get("type")),
			Host: r.PostForm.Get("host"),
			User: r.PostForm.Get("user"),
			Pass: r.PostForm.Get("pass"),
			Name: r.PostForm.Get("name"),
		}
		if conf.Port, err = strconv.ParseUint(r.PostForm.Get("port"), 10, 16); err != nil {
			jRsp(w, http.StatusInternalServerError, err.Error(), nil)
			return
		}
		if err = model.Open(conf); err != nil {
			jRsp(w, http.StatusInternalServerError, err.Error(), nil)
			return
		}
		if err = model.Install(path); err != nil {
			jRsp(w, http.StatusInternalServerError, err.Error(), nil)
			return
		}
		if err = conf.Save(path); err != nil {
			jRsp(w, http.StatusInternalServerError, err.Error(), nil)
			return
		}
		user := &model.Admin{
			Email:    r.PostForm.Get("email"),
			Password: r.PostForm.Get("password"),
			GroupID:  1,
		}
		if err = user.Create(); err != nil {
			jRsp(w, http.StatusInternalServerError, err.Error(), nil)
			return
		}
		Route(route)
		jRsp(w, http.StatusOK, "安装成功", "/login")
	})
}
