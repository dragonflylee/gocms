package handler

import (
	"log"
	"net/http"
	"strings"

	"github.com/dragonflylee/gocms/model"
)

// Home 首页
func Home(w http.ResponseWriter, r *http.Request) {
	render(w, r, "index.tpl", nil)
}

// Profile 个人中心
func Profile(w http.ResponseWriter, r *http.Request) {
	render(w, r, "profile.tpl", nil)
}

// Login 登录页
func Login(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, sessName)
	if err != nil {
		log.Printf("login session err(%s)", err.Error())
	} else if _, exist := session.Values["user"]; exist {
		http.Redirect(w, r, "/admin/", http.StatusFound)
		return
	}
	if r.Method == http.MethodGet {
		t.ExecuteTemplate(w, "login.tpl", nil)
		return
	}
	if err = r.ParseForm(); err != nil {
		rsp(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	user, err := model.Login(
		r.PostForm.Get("username"),
		strings.ToLower(r.PostForm.Get("password")),
		r.RemoteAddr)
	if err != nil {
		rsp(w, http.StatusForbidden, err.Error(), nil)
		return
	}
	session.Values["user"] = user
	session.Save(r, w)
	rsp(w, http.StatusOK, "登录成功", "/admin/")
}

// Logout 登出
func Logout(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, sessName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for key := range session.Values {
		delete(session.Values, key)
	}
	session.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusFound)
}
