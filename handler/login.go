package handler

import (
	"encoding/json"
	"net/http"
	"strings"
	"sync"

	"github.com/dragonflylee/gocms/model"

	"github.com/dchest/captcha"
)

var tokenMap sync.Map

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
	email := strings.TrimSpace(r.PostForm.Get("user"))
	password := strings.TrimSpace(r.PostForm.Get("pass"))
	if !emailRegexp.MatchString(email) {
		jFailed(w, http.StatusBadRequest, "邮箱格式非法")
		return
	}
	if !md5Regexp.MatchString(password) {
		jFailed(w, http.StatusBadRequest, "密码不正确")
		return
	}

	jResp := func(w http.ResponseWriter, msg string) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(struct {
			Code int    `json:"code"`
			Msg  string `json:"msg,omitempty"`
			Data string `json:"data,omitempty"`
		}{Code: http.StatusBadRequest, Msg: msg, Data: captcha.New()})
	}

	if !captcha.VerifyString(r.Form.Get("id"), r.PostForm.Get("code")) {
		jResp(w, "验证码非法")
		return
	}
	user, err := model.Login(email, password, r.RemoteAddr)
	if err != nil {
		jResp(w, err.Error())
		return
	}
	sess.Values[userKey] = user
	if _, exist := r.PostForm["remember"]; !exist {
		sess.Options.MaxAge = 0
	}
	if err = sess.Save(r, w); err != nil {
		jResp(w, err.Error())
		return
	}
	tokenMap.Store(user.ID, sess.ID)
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
