package model

import (
	"encoding/json"
	"gocms/libraries/mongo"
	"gocms/libraries/redis"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// Config 数据库配置项
type Config struct {
	Type      string               `json:"type"`
	Host      string               `json:"host,omitempty"`
	Port      uint64               `json:"port,omitempty"`
	User      string               `json:"user,omitempty"`
	Pass      string               `json:"pass,omitempty"`
	Name      string               `json:"name"`
	RedisConf *redis.RedisConf     `json:"redisConf,omitempty"`
	MongoConf *mongo.MongodbConfig `json:"mongoConf,omitempty"`
}

// Load 加载配置
func (m *Config) Load(path string) error {
	data, err := ioutil.ReadFile(filepath.Join(path, "config.json"))
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
	return ioutil.WriteFile(filepath.Join(path, "config.json"),
		data, os.ModePerm)
}
