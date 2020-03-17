package main

import (
	"flag"
	"gocms/handler"
	"gocms/model"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/dchest/captcha"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/configor"
)

var (
	config = configor.New(&configor.Config{AutoReload: true, AutoReloadInterval: time.Hour})
	path   = flag.String("c", "config.yml", "配置文件路径")
)

func main() {
	flag.Parse()

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// 加载配置文件
	if err := config.Load(&model.Config, *path); err == nil {
		if err = model.Open(config.Debug); err != nil {
			log.Panicf("open db failed: %v", err)
		}
	}
	r := mux.NewRouter()
	// 静态文件
	dir := filepath.Dir(os.Args[0])
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir(filepath.Join(dir, "static")))))
	// 404页面
	r.NotFoundHandler = handler.Limit(2, http.NotFound)
	s := handler.RouterWrap{Router: r.PathPrefix("/").Subrouter()}
	if !model.IsOpen() {
		s.Use(func(h http.Handler) http.Handler {
			if model.IsOpen() {
				return h
			}
			return handler.Install(*path, config.Debug, s.Router)
		})
	}
	s.Use(handlers.ProxyHeaders, handler.LogHandler,
		handlers.RecoveryHandler(handlers.PrintRecoveryStack(true)))
	// 初始化模板
	if err := handler.Watch(filepath.Join(dir, "views"), s.Router); err != nil {
		log.Fatalf("watch failed: %v", err)
	}
	// 登录相关
	s.Handle("/login", handler.Limit(2, handler.Login)).Methods(http.MethodPost)
	s.Handle("/bingpic", handler.Limit(2, handler.BingPic)).Methods(http.MethodGet)
	s.Handle("/captcha/{png}", captcha.Server(120, 35)).Methods(http.MethodGet)

	s.HandleFunc("/login", handler.Login).Methods(http.MethodGet)
	s.HandleFunc("/logout", handler.Logout)
	s.HandleFunc("/password", handler.Password).Methods(http.MethodPost)
	// 后台主页
	s = handler.RouterWrap{Router: s.PathPrefix("/").Subrouter()}
	// 检查登陆状态
	s.Use(handler.Check, handlers.CompressHandler)
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
	s.HandleFunc("/", handler.Home)

	log.Panic(http.ListenAndServe(model.Config.Addr, r))
}
