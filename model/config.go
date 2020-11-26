package model

// Config 数据库配置项
var Config struct {
	Addr string `default:":8081" env:"LISTEN_ADDR"`
	DB   struct {
		Type string `required:"true"`
		Host string `required:"true"`
		Port uint64 `yaml:"port,omitempty"`
		User string `required:"true"`
		Pass string `yaml:"pass,omitempty"`
		Name string `yaml:"name,omitempty"`
	}
	Captcha struct {
		API    string `default:"https://recaptcha.net/recaptcha/api/siteverify"`
		Key    string
		Secret string
	}
}
