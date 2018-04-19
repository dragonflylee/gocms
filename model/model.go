package model

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	db *gorm.DB
)

// Open 连接数据库
func Open(config *Config) (err error) {
	if db, err = gorm.Open("sqlite3", "gocms.db"); err != nil {
		log.Printf("failed to connect database (%s)", err.Error())
		return err
	}
	// 同步数据库
	db.AutoMigrate(&Group{}, &Admin{}, &AdminLog{}, &Node{})
	// 加载节点数据
	if err = initNodes(); err != nil {
		log.Printf("failed init nodes (%s)", err.Error())
		return err
	}
	return nil
}
