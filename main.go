package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/dragonflylee/gocms/handler"
	"github.com/dragonflylee/gocms/model"
	"github.com/jinzhu/configor"

	"github.com/dchest/captcha"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	addr   = flag.String("addr", ":8080", "监听端口")
	config = flag.String("c", "config.yml", "配置文件路径")
)

func main() {
	flag.Parse()

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	path := filepath.Dir(os.Args[0])
	// 初始化模板
	if err := handler.Watch(filepath.Join(path, "views")); err != nil {
		log.Fatalf("watch failed: %v", err)
	}
	r := mux.NewRouter()
	// 静态文件
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir(filepath.Join(path, "static")))))
	// 404页面
	r.NotFoundHandler = handler.Limit(2, http.NotFound)
	// 加载配置文件
	if err := configor.New(nil).Load(&model.Config, *config); err == nil {
		if err = model.Open(); err != nil {
			log.Panicf("open db failed: %v", err)
		}
	}
	s := r.PathPrefix("/").Subrouter()
	if !model.IsOpen() {
		s.Use(func(h http.Handler) http.Handler {
			if model.IsOpen() {
				return h
			}
			return handler.Install(path, s)
		})
	}
	// 登录相关
	s.Handle("/", handler.Check(http.HandlerFunc(handler.Home)))
	s.Handle("/login", handler.Limit(2, handler.Login)).Methods(http.MethodPost)
	s.Handle("/captcha/{png}", captcha.Server(120, 35)).Methods(http.MethodGet)

	s.HandleFunc("/login", handler.Login).Methods(http.MethodGet)
	s.HandleFunc("/logout", handler.Logout)
	s.HandleFunc("/password", handler.Password).Methods(http.MethodPost)
	// 后台主页
	s = s.PathPrefix("/").Subrouter()
	// 检查登陆状态
	s.Use(handler.Check)
	// 系统管理
	s.HandleFunc("/users", handler.Users).Methods(http.MethodGet)
	s.HandleFunc("/user/add", handler.UserAdd).Methods(http.MethodPost)
	s.HandleFunc("/user/delete/{id:[0-9]+}", handler.UserDelete)
	s.HandleFunc("/group/{id:[0-9]+}", handler.GroupEdit)
	s.HandleFunc("/group/add", handler.GroupAdd).Methods(http.MethodPost)
	s.HandleFunc("/logs", handler.Logs).Methods(http.MethodGet)
	// 个人中心
	s.HandleFunc("/profile", handler.Profile).Methods(http.MethodGet)
	// 文件上传
	s.HandleFunc("/upload", handler.Upload).Methods(http.MethodPost)

	log.Panic(http.ListenAndServe(*addr, handlers.CustomLoggingHandler(os.Stdout,
		handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(r), handler.WriteLog)))
}
