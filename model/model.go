package model

import (
	"fmt"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	db       *gorm.DB
	mapNodes map[int]*Node
)

// Open 连接数据库
func Open(debug bool) (err error) {

	config := &gorm.Config{
		AllowGlobalUpdate: false,
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "gc_",
		},
		Logger: logger.Default.LogMode(logger.Info),
	}

	var dialect gorm.Dialector
	switch strings.ToLower(Config.DB.Type) {
	case "mysql":
		dialect = mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&allowOldPasswords=1",
			Config.DB.User, Config.DB.Pass, Config.DB.Host, Config.DB.Port, Config.DB.Name))
	case "postgres":
		dialect = postgres.Open(fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
			Config.DB.User, Config.DB.Pass, Config.DB.Host, Config.DB.Port, Config.DB.Name))
	case "sqlite3":
		dialect = sqlite.Open(fmt.Sprintf("file:%s?_auth&_auth_user=%s&_auth_pass=%s",
			Config.DB.Host, Config.DB.User, Config.DB.Pass))
	default:
		return fmt.Errorf("%s is unsupport", Config.DB.Type)
	}
	if db, err = gorm.Open(dialect, config); err != nil {
		return fmt.Errorf("connect database failed: %v", err)
	}

	// 同步数据库
	if err = db.AutoMigrate(
		new(Admin),
		new(Group),
		new(AdminLog),
		new(AdminRecord),
		new(Node),
		new(Article),
	); err != nil {
		return fmt.Errorf("migrate failed: %v", err)
	}
	// 加载节点数据
	if mapNodes, err = loadNodes(); err != nil {
		return fmt.Errorf("init nodes failed: %v", err)
	}
	if time.Local, err = time.LoadLocation("Asia/Chongqing"); err != nil {
		return fmt.Errorf("load location failed: %v", err)
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
