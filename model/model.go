package model

import (
	"errors"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	// 数据库驱动
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
)

// Open 连接数据库
func Open(conf *Config) error {
	var (
		source string
		err    error
	)
	if conf.Type == "mysql" {
		source = fmt.Sprintf("%s:%s@%s:%d/%s?charset=utf8",
			conf.User, conf.Pass, conf.Host, conf.Port, conf.Name)
	} else if conf.Type == "postgres" {
		source = fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
			conf.User, conf.Pass, conf.Host, conf.Port, conf.Name)
	} else {
		return errors.New("数据库类型不支持")
	}

	if db, err = gorm.Open(conf.Type, source); err != nil {
		log.Printf("failed to connect database (%s)", err.Error())
		return err
	}
	// 同步数据库
	db.AutoMigrate(&Group{}, &Admin{}, &AdminLog{}, &Node{})
	// 加载节点数据
	if err = loadNodes(); err != nil {
		log.Printf("failed init nodes (%s)", err.Error())
		return err
	}
	return nil
}
