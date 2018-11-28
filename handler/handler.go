package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"sync"
	"time"

	"github.com/Tomasen/realip"
	"github.com/dragonflylee/gocms/model"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"golang.org/x/time/rate"
)

const (
	defaultMaxMemory = 32 << 20  // 32 MB
	sessName         = "X-GoCMS" // Session 名称
	dateFormate      = "2006-01-02"
)

var (
	t           = template.New("")
	md5Regexp   = regexp.MustCompile("[a-fA-F0-9]{32}$")
	emailRegexp = regexp.MustCompile("^[a-zA-Z0-9_.-]+@[a-zA-Z0-9-]+(\\.[a-zA-Z0-9-]+)*\\.[a-zA-Z0-9]{2,6}$")
	store       = sessions.NewFilesystemStore(os.TempDir(), securecookie.GenerateRandomKey(32))
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

func jSuccess(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"code": http.StatusOK, "data": data})
}

func jFailed(w http.ResponseWriter, code int64, format string, a ...interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"code": code, "msg": fmt.Sprintf(format, a...)})
}

// rLayout 渲染模板
func rLayout(w http.ResponseWriter, r *http.Request, name string, data interface{}) {
	if session, err := store.Get(r, sessName); err != nil {
		Error(w, http.StatusBadRequest, "页面错误 %v", err)
	} else if s := mux.CurrentRoute(r); s == nil {
		Error(w, http.StatusBadRequest, "页面错误")
	} else if tpl, err := s.GetPathTemplate(); err != nil {
		Error(w, http.StatusNotFound, "页面错误 %v", err)
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

// Error 错误页面
func Error(w http.ResponseWriter, code int, format string, a ...interface{}) {
	w.WriteHeader(code)
	t.ExecuteTemplate(w, "error.tpl", map[string]interface{}{
		"code": code, "text": http.StatusText(code),
		"msg": fmt.Sprintf(format, a...),
	})
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
		} else if token, exist := tokenMap[user.ID]; exist && token != session.ID {
			session.Options.MaxAge = -1
			session.Save(r, w)
			http.Redirect(w, r, "/login", http.StatusFound)
		} else if !user.Status && r.URL.Path != "/profile" {
			http.Redirect(w, r, "/profile", http.StatusFound)
		} else if c := mux.CurrentRoute(r); c == nil {
			Error(w, http.StatusNotFound, "页面错误")
		} else if tpl, err := c.GetPathTemplate(); err != nil {
			Error(w, http.StatusNotFound, "页面错误 %v", err)
		} else if user.Access(tpl) {
			h.ServeHTTP(w, r)
		} else {
			Error(w, http.StatusForbidden, "无权访问 %s", tpl)
		}
	})
}

// Limit 请求限制
func Limit(b int, f func(http.ResponseWriter, *http.Request)) http.Handler {
	var (
		bucket = make(map[string]*rate.Limiter)
		lock   sync.Mutex
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			ip    = realip.FromRequest(r)
			l     *rate.Limiter
			exist bool
		)
		lock.Lock()
		if l, exist = bucket[ip]; !exist {
			l = rate.NewLimiter(rate.Limit(1), b)
			bucket[ip] = l
		}
		lock.Unlock()
		if l.Allow() {
			f(w, r)
			return
		}
		Error(w, http.StatusTooManyRequests, "请求太频繁")
	})
}

// WriteLog 日志打印
func WriteLog(w io.Writer, p handlers.LogFormatterParams) {
	var u string
	if session, err := store.Get(p.Request, sessName); err == nil {
		if cookie, exist := session.Values["user"]; exist {
			u = fmt.Sprint(cookie)
		}
	}
	fmt.Fprintf(w, "%s %s %s %d %d %s %s\n", p.TimeStamp.Format("2006/01/02 15:04:05"), p.Request.Method,
		p.URL.RequestURI(), p.StatusCode, p.Size, realip.FromRequest(p.Request), u)
}

// Start 初始化控制层
func Start(path string) {
	// 注册类型
	pattern := filepath.Join(path, "views", "*.tpl")
	// 注册自定义函数
	t.Funcs(template.FuncMap{
		"date": func(t *time.Time) string {
			if t == nil {
				return "无"
			}
			return t.In(time.Local).Format("2006-01-02 15:04:05")
		},
		"html": func(s string) template.HTML {
			return template.HTML(s)
		},
		"url": func(r url.Values, a ...string) template.URL {
			if len(a) <= 0 {
				return template.URL(r.Encode())
			}
			u := make(url.Values)
			for _, n := range a {
				if v := r.Get(n); len(v) > 0 {
					u.Add(n, v)
				}
			}
			return template.URL(u.Encode())
		},
		"version": func() template.HTML {
			return template.HTML(runtime.Version())
		},
	})
	t = template.Must(t.ParseGlob(pattern))
}
