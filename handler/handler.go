package handler

import (
	"encoding/json"
	"html/template"
	"net/http"
	"path/filepath"
	"runtime"

	"github.com/dragonflylee/gocms/model"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

const (
	defaultMaxMemory = 32 << 20 // 32 MB
	sessName         = "gocms"  // Session 名称
)

var (
	t         = template.New("")
	emptyData = map[string]interface{}{"list": nil, "page": nil}
	store     = sessions.NewFilesystemStore(".", []byte("gocms"))
)

func jRsp(w http.ResponseWriter, code int64, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code": code, "msg": message, "data": data})
}

// render 渲染模板
func rLayout(w http.ResponseWriter, r *http.Request, name string, data interface{}) {
	if session, err := store.Get(r, sessName); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		t.ExecuteTemplate(w, name, map[string]interface{}{
			"menu": model.GetNodes(),
			"node": model.GetNodeByPath(r.URL.Path),
			"user": session.Values["user"],
			"form": r.Form,
			"data": data})
	}
}

func checkLogin(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if session, err := store.Get(r, sessName); err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
		} else if session.IsNew || len(session.Values) < 1 {
			http.Redirect(w, r, "/login", http.StatusFound)
		} else if user, ok := session.Values["user"].(*model.Admin); !ok {
			http.Redirect(w, r, "/login", http.StatusFound)
		} else if c := mux.CurrentRoute(r); c == nil {
			http.NotFound(w, r)
		} else if tpl, err := c.GetPathTemplate(); err != nil {
			http.NotFound(w, r)
		} else if user.Access(tpl) {
			h.ServeHTTP(w, r)
		} else if r.Header.Get("X-Requested-With") != "" {
			jRsp(w, http.StatusForbidden, "无权操作", nil)
		} else {
			http.Error(w, "<h1>Forbidden</h1>", http.StatusForbidden)
		}
	})
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
	r.HandleFunc("/login", Login)
	r.HandleFunc("/logout", Logout)
	// 后台主页
	s := r.PathPrefix("/admin").Subrouter()
	// 检查登陆状态
	s.Use(checkLogin)
	// 系统管理
	s.HandleFunc("/users", Users).Methods(http.MethodGet)
	s.HandleFunc("/logs", Logs).Methods(http.MethodGet)
	// 个人中心
	s.HandleFunc("/profile", Profile).Methods(http.MethodGet)
	s.HandleFunc("", Home).Methods(http.MethodGet)
}
