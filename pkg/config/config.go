package config

import (
	"os"
	"time"

	"github.com/jinzhu/configor"
	"gopkg.in/yaml.v2"
)

// Config 配置项
var (
	config struct {
		HTTP    HttpConfig
		Session SessionConfig
		UI      UIConfig
		DB      DBConfig
		Captcha *CaptchaConfig
		SMTP    *SmtpConfig
	}

	callback = make([]func(), 0)
	configer = configor.New(&configor.Config{
		AutoReload:         true,
		AutoReloadInterval: time.Hour,
		AutoReloadCallback: func(config interface{}) {
			for _, f := range callback {
				f()
			}
		},
	})
)

type HttpConfig struct {
	Addr string `env:"HTTP_ADDR" default:"127.0.0.1:8081"`
}

type SessionConfig struct {
	Type   string `env:"SESSION_TYPE" default:"file"`
	Name   string `env:"SESSION_NAME" default:"X-GoCMS"`
	Secret string `env:"SESSION_SECRET"`
}

type UIConfig struct {
	Theme string `env:"UI_THEME" default:"adminlte"`
	Lang  string `env:"Ui_LANG" default:"zh-CN"`
}

type DBConfig struct {
	Type   string `env:"DB_TYPE" form:"db_type" binding:"required" required:"true"`
	Host   string `env:"DB_HOST" form:"db_host" binding:"required" required:"true"`
	User   string `env:"DB_USER" form:"db_user" binding:"required" required:"true"`
	Pass   string `env:"DB_PASS" form:"db_pass" binding:"required"`
	Name   string `env:"DB_NAME" form:"db_name" binding:"required"`
	Prefix string `env:"DB_PREFIX" form:"db_prefix"`
}

type CaptchaConfig struct {
	API    string `env:"CAPTCHA_API" default:"https://recaptcha.net/recaptcha/api/siteverify"`
	Key    string `env:"CAPTCHA_KEY"`
	Secret string `env:"CAPTCHA_SECRET"`
}

type SmtpConfig struct {
	Host string `env:"SMTP_HOST"`
	Port int    `env:"SMTP_PORT"`
	User string `env:"SMTP_USER"`
	Pass string `env:"SMTP_PASS"`
	SSL  bool   `env:"SMTP_SSL"`
}

// Load 加载配置文件
func Load(path string) error {
	if err := configer.Load(&config, path); err != nil {
		return err
	}
	configer.AutoReloadCallback(&config)
	return nil
}

// Save 保存配置文件
func Save(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return yaml.NewEncoder(f).Encode(&config)
}

// Register 注册配置加载回调
func Register(f func()) {
	callback = append(callback, f)
}

func Debug() bool {
	return configer.Debug
}

func HTTP() *HttpConfig {
	return &config.HTTP
}

func Session() *SessionConfig {
	return &config.Session
}

func UI() *UIConfig {
	return &config.UI
}

func DB() *DBConfig {
	return &config.DB
}

func Captcha() *CaptchaConfig {
	return config.Captcha
}

func Smtp() *SmtpConfig {
	return config.SMTP
}
