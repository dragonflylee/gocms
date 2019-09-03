package handler

import (
	"net/http"
	"strings"
	"sync"

	"github.com/dragonflylee/gocms/model"

	"github.com/Tomasen/realip"
	"github.com/dchest/captcha"
)

var (
	tokenMap   = make(map[int64]string)
	tokenMutex sync.RWMutex
)

// Login 登录页
func Login(w http.ResponseWriter, r *http.Request) {
	sess, err := store.Get(r, sessName)
	if err == nil {
		if _, exist := sess.Values[userKey]; exist {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
	}
	if r.Method == http.MethodGet {
		t.ExecuteTemplate(w, "login.tpl", map[string]string{
			"Ref": r.Referer(), "Captcha": captcha.New()})
		return
	}
	if err = r.ParseForm(); err != nil {
		jFailed(w, http.StatusBadRequest, err.Error())
		return
	}
	if !captcha.VerifyString(r.PostForm.Get("id"), r.PostForm.Get("code")) {
		jFailed(w, http.StatusBadRequest, "验证码非法")
		return
	}
	email := strings.TrimSpace(r.PostForm.Get("username"))
	password := strings.TrimSpace(r.PostForm.Get("password"))
	if !emailRegexp.MatchString(email) {
		jFailed(w, http.StatusBadRequest, "邮箱格式非法")
		return
	}
	if !md5Regexp.MatchString(password) {
		jFailed(w, http.StatusBadRequest, "密码不正确")
		return
	}
	user, err := model.Login(email, password, realip.FromRequest(r))
	if err != nil {
		jFailed(w, http.StatusForbidden, err.Error())
		return
	}
	sess.Values[userKey] = user
	if _, exist := r.PostForm["remember"]; !exist {
		sess.Options.MaxAge = 3600
	}
	if err = sess.Save(r, w); err != nil {
		jFailed(w, http.StatusForbidden, err.Error())
		return
	}
	tokenMutex.Lock()
	defer tokenMutex.Unlock()
	tokenMap[user.ID] = sess.ID
	jSuccess(w, r.Form.Get("refer"))
}

// Logout 登出
func Logout(w http.ResponseWriter, r *http.Request) {
	if sess, err := store.Get(r, sessName); err != nil {
		Error(w, http.StatusBadRequest, err.Error())
	} else {
		sess.Options.MaxAge = -1
		sess.Save(r, w)
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}
