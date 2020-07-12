package handler

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"gocms/model"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/quasoft/memstore"
	"golang.org/x/time/rate"
)

type sessKey int

const (
	userKey  sessKey = 1
	sessName         = "X-GoCMS" // session 名称

	defaultMaxMemory = 32 << 20 // 32 MB
	dateFormate      = "2006-01-02"
)

var (
	t *template.Template

	build = "0"
	store = memstore.NewMemStore(securecookie.GenerateRandomKey(32))

	emailRegexp = regexp.MustCompile("^[a-zA-Z0-9_.-]+@[a-zA-Z0-9-]+(\\.[a-zA-Z0-9-]+)*\\.[a-zA-Z0-9]{2,6}$")
)

func aLog(r *http.Request, format string, a ...interface{}) error {
	m := &model.AdminLog{
		Path:   r.URL.String(),
		UA:     r.UserAgent(),
		IP:     r.RemoteAddr,
		Commit: fmt.Sprintf(format, a...),
	}
	if sess, err := store.Get(r, sessName); err != nil {
		return err
	} else if cookie, exist := sess.Values[userKey]; !exist {
		return http.ErrNoCookie
	} else if user, ok := cookie.(*model.Admin); ok {
		m.AdminID = user.ID
	}
	return m.Create()
}

func jSuccess(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Code int         `json:"code"`
		Data interface{} `json:"data,omitempty"`
	}{Code: http.StatusOK, Data: data})
}

func jFailed(w http.ResponseWriter, code int, format string, a ...interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Code int    `json:"code"`
		Msg  string `json:"msg,omitempty"`
	}{Code: code, Msg: fmt.Sprintf(format, a...)})
}

// rLayout 渲染模板
func rLayout(w http.ResponseWriter, r *http.Request, name string, data interface{}) {
	if sess, err := store.Get(r, sessName); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if s := mux.CurrentRoute(r); s == nil {
		http.Error(w, "BadRequest", http.StatusBadRequest)
	} else if err = t.ExecuteTemplate(w, name, map[string]interface{}{
		"Menu": model.GetNodes(),
		"Node": model.GetNodeByPath(s.GetName()),
		"User": sess.Values[userKey],
		"Form": r.Form, "Data": data,
	}); err != nil {
		fmt.Fprint(w, err)
	}
}

// Check 检查用户登录
func Check(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if sess, err := store.Get(r, sessName); err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
		} else if cookie, exist := sess.Values[userKey]; !exist {
			http.Redirect(w, r, "/login", http.StatusFound)
		} else if user, ok := cookie.(*model.Admin); !ok {
			http.Redirect(w, r, "/login", http.StatusFound)
		} else if token, exist := tokenMap.Load(user.ID); exist && token != sess.ID {
			sess.Options.MaxAge = -1
			sess.Save(r, w)
			http.Redirect(w, r, "/login", http.StatusFound)
		} else if !user.Status() && r.URL.Path != "/profile" {
			http.Redirect(w, r, "/profile", http.StatusFound)
		} else if c := mux.CurrentRoute(r); c == nil {
			http.Error(w, "BadRequest", http.StatusBadRequest)
		} else if user.Access(c.GetName()) {
			h.ServeHTTP(w, r)
		} else {
			http.Error(w, r.URL.RawPath, http.StatusForbidden)
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
		lock.Lock()
		l, exist := bucket[r.RemoteAddr]
		if !exist {
			l = rate.NewLimiter(rate.Limit(1), b)
			bucket[r.RemoteAddr] = l
		}
		lock.Unlock()
		if l.Allow() {
			f(w, r)
			return
		}
		http.Error(w, "TooManyRequests", http.StatusTooManyRequests)
	})
}

// LogHandler 日志打印
func LogHandler(h http.Handler) http.Handler {
	return handlers.CustomLoggingHandler(os.Stdout, h, func(w io.Writer, p handlers.LogFormatterParams) {
		var u string
		if sess, err := store.Get(p.Request, sessName); err == nil {
			if cookie, exist := sess.Values[userKey]; exist {
				u = fmt.Sprint(cookie)
			}
		}
		fmt.Fprintf(w, "%s %s %d %s %d %s (%s) %s\n", p.TimeStamp.Format("2006/01/02 15:04:05"),
			p.Request.Method, p.StatusCode, p.URL.RequestURI(), p.Size,
			p.Request.RemoteAddr, time.Since(p.TimeStamp), u)
	})
}

// RouterWrap 路由封装
type RouterWrap struct {
	*mux.Router
}

// HandleFunc 反射函数名
func (r RouterWrap) HandleFunc(path string, h http.HandlerFunc) *mux.Route {
	n := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
	return r.Handle(path, h).Name(n[strings.LastIndexByte(n, '.')+1:])
}

// Watch 初始化控制层
func Watch(tpl string, r *mux.Router) error {
	// 注册自定义函数
	funcMap := template.FuncMap{
		"date": func(t *time.Time) string {
			if t == nil {
				return "-"
			}
			return t.In(time.Local).Format("2006-01-02 15:04:05")
		},
		"urlfor": func(name string, pair ...string) (*url.URL, error) {
			if s := r.Get(name); s != nil {
				return s.URL(pair...)
			}
			return &url.URL{Fragment: "top"}, nil
		},
		"html": func(s string) template.HTML {
			return template.HTML(s)
		},
		"version": func() string {
			return fmt.Sprintf("1.0.%s", build)
		},
	}
	// 文件监控
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return fmt.Errorf("create watcher: %v", err)
	}
	go func() {
		pattern := filepath.Join(tpl, "*.tpl")
		for {
			select {
			case e := <-watcher.Events:
				log.Printf("load %s %v", filepath.Base(e.Name), e.Op)
				if t, err = template.New(sessName).Funcs(funcMap).ParseGlob(pattern); err != nil {
					log.Printf("parse %s failed: %v", e.Name, err)
				}
			case err := <-watcher.Errors:
				log.Printf("Watcher error: %v", err) // No need to exit here
			}
		}
	}()
	watcher.Events <- fsnotify.Event{}
	return watcher.Add(tpl)
}

func init() {
	gob.Register(userKey)
}
