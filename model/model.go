package model

import (
	"errors"
	"flag"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	// 数据库驱动
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db       *gorm.DB
	mapNodes map[int64]*Node
	debug    = flag.Bool("d", false, "debug mode")
)

// Open 连接数据库
func Open(conf *Config) error {
	var (
		source string
		err    error
	)
	if conf.Type == "mysql" {
		source = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&allowOldPasswords=1",
			conf.User, conf.Pass, conf.Host, conf.Port, conf.Name)
	} else if conf.Type == "postgres" {
		source = fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
			conf.User, conf.Pass, conf.Host, conf.Port, conf.Name)
	} else {
		return errors.New("数据库类型不支持")
	}
	if db, err = gorm.Open(conf.Type, source); err != nil {
		return fmt.Errorf("connect database failed: %v", err)
	}
	db.BlockGlobalUpdate(true)
	if debug != nil {
		db.LogMode(*debug)
	}
	// 同步数据库
	if err = db.AutoMigrate(&Group{}, &Admin{}, &AdminLog{}, &Node{}).Error; err != nil {
		return fmt.Errorf("migrate failed: %v", err)
	}
	// 加载节点数据
	if mapNodes, err = loadNodes(); err != nil {
		return fmt.Errorf("init nodes failed: %v", err)
	}
	if time.Local, err = time.LoadLocation("Asia/Chongqing"); err != nil {
		return fmt.Errorf("load location failed: %v", err)
	}
	gorm.NowFunc = func() time.Time {
		return time.Now().UTC()
	}
	return nil
}

// IsOpen 数据库是否连接
func IsOpen() bool {
	if db == nil {
		return false
	}
	if db.Error != nil {
		return false
	}
	if mapNodes == nil {
		return false
	}
	return true
}
