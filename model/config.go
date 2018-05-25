package model

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

// Config 数据库配置项
type Config struct {
	Type string `json:"type"`
	Host string `json:"host,omitempty"`
	Port uint64 `json:"port,omitempty"`
	User string `json:"user,omitempty"`
	Pass string `json:"pass,omitempty"`
	Name string `json:"name"`
}

// Load 加载配置
func (m *Config) Load(path string) error {
	f, err := os.Open(filepath.Join(path, "config.json"))
	if err != nil {
		log.Printf("failed load config (%s)", err.Error())
		return err
	}
	defer f.Close()
	return json.NewDecoder(f).Decode(m)
}

// Save 保存配置
func (m *Config) Save(path string) error {
	f, err := os.Create(filepath.Join(path, "config.json"))
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(m)
}
