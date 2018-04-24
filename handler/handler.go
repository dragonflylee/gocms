package handler

import (
	"encoding/json"
	"html/template"
	"math/rand"
	"net/http"
	"path/filepath"
	"runtime"
	"time"

	"github.com/dragonflylee/gocms/model"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

const (
	defaultMaxMemory = 32 << 20 // 32 MB
	sessName         = "gocms"  // Session 名称
	indexPath        = "/admin"
	loginPath        = "/login"
)

var (
	t         = template.New("")
	emptyData = map[string]interface{}{"list": nil, "page": nil}
	store     = sessions.NewFilesystemStore(".", []byte(randString(15)))
)

func jRsp(w http.ResponseWriter, code int64, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code": code, "msg": message, "data": data})
}

// render 渲染模板
func rLayout(w http.ResponseWriter, r *http.Request, name string, data interface{}) {
	session, err := store.Get(r, sessName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.ExecuteTemplate(w, name, map[string]interface{}{
		"menu": model.GetNodes(),
		"node": model.GetNodeByPath(r.URL.Path),
		"user": session.Values["user"],
		"form": r.Form,
		"data": data})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// randString 生成随机字符串
func randString(l int) string {
	bytes := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// Start 初始化控制层
func Start(path string) {
	// 注册类型
	pattern := filepath.Join(path, "views", "*.tpl")
	// 注册自定义函数
	t.Funcs(template.FuncMap{
		"html": func(s string) template.HTML {
			return template.HTML(s)
		},
		"version": func() template.HTML {
			return template.HTML(runtime.Version())
		},
	})
	t = template.Must(t.ParseGlob(pattern))
}

// Route 初始化路由
func Route(r *mux.Router) {
	if t := r.Get("index"); t != nil {
		t.HandlerFunc(Login)
	} else {
		r.HandleFunc("/", Login)
	}
	// 登录相关
	r.HandleFunc(loginPath, Login)
	r.HandleFunc("/logout", Logout)
	// 后台主页
	s := r.PathPrefix(indexPath).Subrouter()
	s.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if session, err := store.Get(r, sessName); err != nil {
				http.Redirect(w, r, loginPath, http.StatusFound)
			} else if session.IsNew || len(session.Values) < 1 {
				http.Redirect(w, r, loginPath, http.StatusFound)
			} else {
				h.ServeHTTP(w, r)
			}
		})
	})
	// 个人中心
	s.HandleFunc("/profile", Profile)
	s.HandleFunc("", Home).Methods(http.MethodGet)
}
