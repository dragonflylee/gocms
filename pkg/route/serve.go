package route

import (
	"context"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"gocms/pkg/config"
	"gocms/pkg/i18n"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
)

func New(fsys fs.FS) *gin.Engine {
	r := gin.New()
	r.Use(gin.LoggerWithFormatter(LogFormat), i18n.Handle, gin.Recovery())
	// 注册自定义函数
	r.SetFuncMap(template.FuncMap{
		"date": func(t *time.Time) string {
			if t == nil {
				return "-"
			}
			return t.In(time.Local).Format("2006-01-02 15:04:05")
		},
		"html": func(s string) template.HTML {
			return template.HTML(s)
		},
		"version": func() string {
			return fmt.Sprintf("1.0.%s", "0")
		},
	})

	// 初始化模板
	pattern := "views/admin/*.html"
	t, err := template.New("").Funcs(r.FuncMap).ParseFS(fsys, pattern)
	if err != nil {
		log.Fatalf("Parse template %v", err)
	}
	r.SetHTMLTemplate(t)

	if config.Debug() {
		go Watch(r.LoadHTMLGlob, pattern)
	}

	// 静态文件
	r.GET("/assets/*filepath", func(c *gin.Context) {
		c.Header("Cache-Control", "public, max-age=3600")
		c.FileFromFS(c.Request.URL.Path, http.FS(fsys))
	})

	return r
}

// Listen HTTP服务优雅退出
func Listen(srv *http.Server, f ...func()) <-chan struct{} {
	idleConnsClosed := make(chan struct{})

	go func() {
		// 监听进程退出
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGTERM, syscall.SIGINT)

		<-sigint
		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	for _, v := range f {
		srv.RegisterOnShutdown(v)
	}

	log.Printf("Listening and serving HTTP on %s", srv.Addr)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}
	return idleConnsClosed
}

// Watch 初始化控制层
func Watch(loadGlob func(string), patterns string) error {
	// 文件监控
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}

	name := filepath.Dir(patterns)
	if err = watcher.Add(name); err != nil {
		return err
	}
	log.Printf("watching %s", name)

	for {
		select {
		case e := <-watcher.Events:
			log.Printf("load %s %v", filepath.Base(e.Name), e.Op)
			loadGlob(patterns)
		case err := <-watcher.Errors:
			log.Printf("Watcher error: %v", err) // No need to exit here
		}
	}
}
