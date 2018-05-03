package handler

import (
	"encoding/json"
	"html/template"
	"net/http"
	"path/filepath"
	"runtime"
	"time"

	"github.com/Tomasen/realip"
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

func aLog(r *http.Request, message string) error {
	var log model.AdminLog
	if session, err := store.Get(r, sessName); err != nil {
		return err
	} else if user, ok := session.Values["user"].(*model.Admin); ok {
		log.AdminID = user.ID
	}
	log.Path = r.URL.String()
	log.UA = r.UserAgent()
	log.IP = realip.FromRequest(r)
	log.Commit = message
	return log.Create()
}

func jRsp(w http.ResponseWriter, code int64, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code": code, "msg": message, "data": data})
}

// rLayout 渲染模板
func rLayout(w http.ResponseWriter, r *http.Request, name string, data interface{}) {
	if session, err := store.Get(r, sessName); err != nil {
		http.NotFound(w, r)
	} else if s := mux.CurrentRoute(r); s == nil {
		http.NotFound(w, r)
	} else if tpl, err := s.GetPathTemplate(); err != nil {
		http.NotFound(w, r)
	} else if err = t.ExecuteTemplate(w, name, map[string]interface{}{
		"menu": model.GetNodes(),
		"node": model.GetNodeByPath(tpl),
		"user": session.Values["user"],
		"form": r.Form,
		"data": data,
	}); err != nil {
		w.Write([]byte(err.Error()))
	}
}

// Check 检查用户登录
func Check(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if session, err := store.Get(r, sessName); err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
		} else if cookie, exist := session.Values["user"]; !exist {
			http.Redirect(w, r, "/login", http.StatusFound)
		} else if user, ok := cookie.(*model.Admin); !ok {
			http.Redirect(w, r, "/login", http.StatusFound)
		} else if !user.Status && r.URL.Path != "/admin/profile" {
			http.Redirect(w, r, "/admin/profile", http.StatusFound)
		} else if c := mux.CurrentRoute(r); c == nil {
			http.NotFound(w, r)
		} else if tpl, err := c.GetPathTemplate(); err != nil {
			http.NotFound(w, r)
		} else if user.Access(tpl) {
			h.ServeHTTP(w, r)
		} else if r.Header.Get("X-Requested-With") != "" {
			jRsp(w, http.StatusForbidden, "无权操作", nil)
		} else {
			http.NotFound(w, r)
		}
	})
}

// Start 初始化控制层
func Start(path string) {
	// 注册类型
	pattern := filepath.Join(path, "views", "*.tpl")
	// 注册自定义函数
	t.Funcs(template.FuncMap{
		"date": func(t time.Time) string {
			return t.Format("2006-01-02 15:04:05")
		},
		"html": func(s string) template.HTML {
			return template.HTML(s)
		},
		"version": func() template.HTML {
			return template.HTML(runtime.Version())
		},
	})
	t = template.Must(t.ParseGlob(pattern))
}
