package model

// Config 数据库配置项
var Config struct {
	DB struct {
		Type string `required:"true"`
		Host string
		Port uint64
		User string
		Pass string
		Name string
	} `required:"true"`
}
