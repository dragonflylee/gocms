package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Tomasen/realip"
	"github.com/dragonflylee/gocms/model"
	"github.com/gorilla/mux"
)

// Install 安装配置
func Install(path string, s *mux.Router) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			t.ExecuteTemplate(w, "install.tpl", nil)
			return
		}
		err := r.ParseForm()
		if err != nil {
			jFailed(w, http.StatusInternalServerError, err.Error())
			return
		}
		user := &model.Admin{
			Email:    strings.ToLower(strings.TrimSpace(r.PostForm.Get("email"))),
			Password: strings.TrimSpace(r.PostForm.Get("password")),
			Headpic:  "/static/img/avatar.png",
			Group:    model.Group{Name: "超级管理员"},
			LastIP:   realip.FromRequest(r),
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
		conf := &model.Config{
			Type: strings.ToLower(r.PostForm.Get("type")),
			Host: r.PostForm.Get("host"),
			User: r.PostForm.Get("user"),
			Pass: r.PostForm.Get("pass"),
			Name: r.PostForm.Get("name"),
		}
		if conf.Port, err = strconv.ParseUint(r.PostForm.Get("port"), 10, 16); err != nil {
			jFailed(w, http.StatusInternalServerError, err.Error())
			return
		}
		if err = model.Open(conf); err != nil {
			jFailed(w, http.StatusInternalServerError, err.Error())
			return
		}
		if err = model.Install(user, path); err != nil {
			jFailed(w, http.StatusInternalServerError, err.Error())
			return
		}
		if err = conf.Save(path); err != nil {
			jFailed(w, http.StatusInternalServerError, err.Error())
			return
		}
		jSuccess(w, "/login")
	})
}
