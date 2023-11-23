package model

import (
	"fmt"
	"strings"
	"time"

	"gocms/pkg/config"

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
func Open() (err error) {
	var dialect gorm.Dialector
	conf := config.DB()

	switch strings.ToLower(conf.Type) {
	case "mysql":
		dialect = mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&allowOldPasswords=1",
			conf.User, conf.Pass, conf.Host, conf.Port, conf.Name))
	case "postgres":
		dialect = postgres.Open(fmt.Sprintf("user=%s password=%s host=%s port=%d database=%s sslmode=disable",
			conf.User, conf.Pass, conf.Host, conf.Port, conf.Name))
	case "sqlite3":
		dialect = sqlite.Open(fmt.Sprintf("file:%s?_auth&_auth_user=%s&_auth_pass=%s",
			conf.Host, conf.User, conf.Pass))
	default:
		return fmt.Errorf("%s is unsupport", conf.Type)
	}

	opt := &gorm.Config{
		AllowGlobalUpdate: false,
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: conf.Prefix,
		},
	}

	if config.Debug() {
		opt.Logger = logger.Default.LogMode(logger.Info)
	}

	if db, err = gorm.Open(dialect, opt); err != nil {
		return fmt.Errorf("connect database failed: %w", err)
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
		return fmt.Errorf("migrate failed: %w", err)
	}
	// 加载节点数据
	if mapNodes, err = loadNodes(); err != nil {
		return fmt.Errorf("init nodes failed: %w", err)
	}
	if time.Local, err = time.LoadLocation("Asia/Chongqing"); err != nil {
		return fmt.Errorf("load location failed: %w", err)
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
