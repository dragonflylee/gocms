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
	if err = conf.Load(path); err != nil {
		r.Handle("/", handler.Install(path, r)).Name("index")
	} else {
		model.Open(&conf)
		handler.Route(r)
	}
	log.Panic(http.ListenAndServe(*addr, r))
}
