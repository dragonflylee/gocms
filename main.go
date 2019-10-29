package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/dchest/captcha"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/configor"

	"github.com/dragonflylee/gocms/handler"
	"github.com/dragonflylee/gocms/model"
)

var (
	addr   = flag.String("addr", ":8080", "监听端口")
	config = flag.String("c", "config.yml", "配置文件路径")
)

func main() {
	flag.Parse()

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	r := mux.NewRouter()
	// 静态文件
	path := filepath.Dir(os.Args[0])
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
	s := handler.RouterWrap{Router: r.PathPrefix("/").Subrouter()}
	if !model.IsOpen() {
		s.Use(func(h http.Handler) http.Handler {
			if model.IsOpen() {
				return h
			}
			return handler.Install(*config, s.Router)
		})
	}
	s.Use(handler.LogHandler, handlers.RecoveryHandler(handlers.PrintRecoveryStack(true)))
	// 初始化模板
	if err := handler.Watch(filepath.Join(path, "views"), s.Router); err != nil {
		log.Fatalf("watch failed: %v", err)
	}
	// 登录相关
	s.Handle("/", handler.Check(http.HandlerFunc(handler.Home))).Name("")
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

	log.Panic(http.ListenAndServe(*addr, r))
}
