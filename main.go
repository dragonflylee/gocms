package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/dragonflylee/gocms/handler"
	"github.com/dragonflylee/gocms/model"
	"github.com/gorilla/mux"
)

var (
	conf model.Config
	r    = mux.NewRouter()
	addr = flag.String("addr", ":8080", "server listen address")
)

func main() {
	flag.Parse()

	path, err := exec.LookPath(os.Args[0])
	if err != nil {
		log.Panicf("gocms service path (%s)", err.Error())
	}
	path = filepath.Dir(path)
	// 初始化控制器
	log.Printf("gocms starting from (%s)", path)
	// 初始化模板
	handler.Start(path)
	// 静态文件
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir(filepath.Join(path, "static")))))
	// 加载配置文件
	if err = conf.Load(path); err == nil {
		model.Open(&conf)
	}
	r.Use(func(h http.Handler) http.Handler {
		if model.IsOpen() {
			return h
		}
		return handler.Install(path, r)
	})
	// 登录相关
	r.HandleFunc("/", handler.Login)
	r.HandleFunc("/login", handler.Login)
	r.HandleFunc("/logout", handler.Logout)
	r.HandleFunc("/password", handler.Password).Methods(http.MethodPost)
	// 后台主页
	s := r.PathPrefix("/admin").Subrouter()
	// 检查登陆状态
	s.Use(handler.Check)
	// 系统管理
	s.HandleFunc("/users", handler.Users).Methods(http.MethodGet)
	s.HandleFunc("/user/add", handler.UserAdd).Methods(http.MethodPost)
	s.HandleFunc("/logs", handler.Logs).Methods(http.MethodGet)
	// 个人中心
	s.HandleFunc("/profile", handler.Profile).Methods(http.MethodGet)
	s.HandleFunc("", handler.Home).Methods(http.MethodGet)

	log.Panic(http.ListenAndServe(*addr, r))
}
