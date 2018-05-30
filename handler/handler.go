package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/Tomasen/realip"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"gocms/model"
)

const (
	defaultMaxMemory = 32 << 20 // 32 MB
	sessName         = "gocms"  // Session 名称
	dateFormate      = "2006-01-02"
)

type select2 struct {
	ID   string `json:"id"`
	Name string `json:"text"`
}

var (
	t     = template.New("")
	store = sessions.NewFilesystemStore(os.TempDir(), securecookie.GenerateRandomKey(32))
)

func aLog(r *http.Request, format string, a ...interface{}) error {
	m := &model.AdminLog{
		Path:   r.URL.String(),
		UA:     r.UserAgent(),
		IP:     realip.FromRequest(r),
		Commit: fmt.Sprintf(format, a...),
	}
	if session, err := store.Get(r, sessName); err != nil {
		return err
	} else if cookie, exist := session.Values["user"]; !exist {
		return http.ErrNoCookie
	} else if user, ok := cookie.(*model.Admin); ok {
		m.AdminID = user.ID
	}
	return m.Create()
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
		"rate": func(r int64) string {
			if r == 0 {
				return "-"
			}
			return fmt.Sprintf("%.2f%%", float64(r)/100)
		},
	})
	t = template.Must(t.ParseGlob(pattern))
}
