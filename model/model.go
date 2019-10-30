package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	// 数据库驱动
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	db       *gorm.DB
	model    []interface{}
	mapNodes map[int64]*Node
)

// Open 连接数据库
func Open(debug bool) (err error) {
	var dsn string
	dialect := strings.ToLower(Config.DB.Type)
	switch dialect {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&allowOldPasswords=1",
			Config.DB.User, Config.DB.Pass, Config.DB.Host, Config.DB.Port, Config.DB.Name)
	case "postgres":
		dsn = fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
			Config.DB.User, Config.DB.Pass, Config.DB.Host, Config.DB.Port, Config.DB.Name)
	case "sqlite3":
		dsn = fmt.Sprintf("file:%s?_auth&_auth_user=%s&_auth_pass=%s",
			Config.DB.Host, Config.DB.User, Config.DB.Pass)
	default:
		return fmt.Errorf("%s is unsupport", Config.DB.Type)
	}
	if db, err = gorm.Open(dialect, dsn); err != nil {
		return fmt.Errorf("connect database failed: %v", err)
	}
	db.BlockGlobalUpdate(true)
	db.LogMode(debug)
	// 同步数据库
	if err = db.AutoMigrate(model...).Error; err != nil {
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
