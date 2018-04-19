package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// Config 数据库配置项
type Config struct {
	Host string `json:"host,omitempty"`
	Port int    `json:"port,omitempty"`
	User string `json:"user,omitempty"`
	Pass string `json:"pass,omitempty"`
	Name string `json:"name"`
}

// Load 加载配置
func (m *Config) Load(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("failed load config (%s)", err.Error())
		return err
	}
	if err = json.Unmarshal(data, m); err != nil {
		log.Printf("failed parse config (%s)", err.Error())
		return err
	}
	return nil
}

// Save 保存配置
func (m *Config) Save(path string) error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, data, os.ModeType)
}
