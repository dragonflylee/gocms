package handler

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"gocms/model"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

var (
	tokenMap sync.Map

	tokenCache = cache.New(time.Minute*5, time.Hour)
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
			"Ref": r.Referer(), "Key": model.Config.Captcha.Key,
		})
		return
	}
	var p *model.AdminLogin

	switch r.FormValue("action") {
	case "verfiy":
		// 二步验证
		token := strings.TrimSpace(r.PostForm.Get("token"))
		if v, ok := tokenCache.Get(token); !ok {
			jFailed(w, http.StatusForbidden, "非法请求")
			return
		} else if p, ok = v.(*model.AdminLogin); !ok {
			jFailed(w, http.StatusForbidden, "非法请求")
			return
		}
		if p.IP != r.RemoteAddr {
			jFailed(w, http.StatusForbidden, "非法请求")
			return
		}
		code := strings.TrimSpace(r.PostForm.Get("code"))
		if len(code) < 1 {
			jFailed(w, http.StatusForbidden, "验证码非法")
			return
		}
		p.Verifyed = true
		// 移除 token
		tokenCache.Delete(token)

	default:
		p = &model.AdminLogin{
			Email:    strings.TrimSpace(r.PostForm.Get("username")),
			Password: strings.TrimSpace(r.PostForm.Get("password")),
			IP:       r.RemoteAddr, UA: r.UserAgent(),
		}
		if p.Email = strings.ToLower(p.Email); !emailRegexp.MatchString(p.Email) {
			jFailed(w, http.StatusBadRequest, "邮箱格式非法")
			return
		}
		if err = model.Recaptcha(r); err != nil {
			jFailed(w, http.StatusBadRequest, "验证码非法")
			return
		}
	}
	u, err := p.Login()
	if err != nil {
		jFailed(w, http.StatusBadRequest, err.Error())
		return
	}
	if err == nil {
		sess.Values[userKey] = u
		if _, exist := r.PostForm["remember"]; !exist {
			sess.Options.MaxAge = 0
		}
		if err = sess.Save(r, w); err != nil {
			jFailed(w, http.StatusBadRequest, err.Error())
			return
		}
		tokenMap.Store(u.ID, sess.ID)
		jSuccess(w, r.Form.Get("refer"))
		return
	}
	if u != nil {
		hash := hmac.New(sha256.New, net.ParseIP(r.RemoteAddr))
		fmt.Fprint(hash, u.ID, r.UserAgent(), time.Now().Unix())
		token := base64.StdEncoding.EncodeToString(hash.Sum(nil))
		tokenCache.SetDefault(token, p)

		w.Header().Add("X-Form-Action", "?action=verfiy")

		buff := &bytes.Buffer{}
		t.ExecuteTemplate(buff, "login-verify", map[string]string{
			"Token": token, "Email": p.Email,
		})
		jSuccess(w, buff.String())
		return
	}
	jFailed(w, http.StatusForbidden, err.Error())
}

// Logout 登出
func Logout(w http.ResponseWriter, r *http.Request) {
	if sess, err := store.Get(r, sessName); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		sess.Options.MaxAge = -1
		sess.Save(r, w)
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}
