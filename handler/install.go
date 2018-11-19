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
			jRsp(w, http.StatusInternalServerError, err.Error(), nil)
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
			jRsp(w, http.StatusBadRequest, "邮箱格式非法", nil)
			return
		}
		if !md5Regexp.MatchString(user.Password) {
			jRsp(w, http.StatusBadRequest, "密码不正确", nil)
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
		if err = model.Install(user, path); err != nil {
			jRsp(w, http.StatusInternalServerError, err.Error(), nil)
			return
		}
		if err = conf.Save(path); err != nil {
			jRsp(w, http.StatusInternalServerError, err.Error(), nil)
			return
		}
		jRsp(w, http.StatusOK, "安装成功", "/login")
	})
}
