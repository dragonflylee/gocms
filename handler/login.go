package handler

import (
	"log"
	"net/http"
	"strings"

	"github.com/Tomasen/realip"
	"github.com/dragonflylee/gocms/model"
)

// Login 登录页
func Login(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, sessName)
	if err != nil {
		log.Printf("login session err(%s)", err.Error())
	} else if _, exist := session.Values["user"]; exist {
		http.Redirect(w, r, "/admin", http.StatusFound)
		return
	}
	if r.Method == http.MethodGet {
		t.ExecuteTemplate(w, "login.tpl", nil)
		return
	}
	if err = r.ParseForm(); err != nil {
		jRsp(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	user, err := model.Login(
		strings.ToLower(r.PostForm.Get("username")),
		strings.ToLower(r.PostForm.Get("password")),
		realip.FromRequest(r))
	if err != nil {
		jRsp(w, http.StatusForbidden, err.Error(), nil)
		return
	}
	session.Values["user"] = user
	session.Save(r, w)
	jRsp(w, http.StatusOK, "登录成功", "/admin")
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
