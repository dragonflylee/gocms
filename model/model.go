package model

import (
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"gocms/libraries/mongo"
	"gocms/libraries/redis"
	mongodb "gopkg.in/mgo.v2"
	// 数据库驱动
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db       *gorm.DB
	mapNodes map[int64]*Node
	debug     = flag.Bool("d", false, "debug mode")
	redisPool *redis.RedisPool
	mgo       *mongodb.Session
	mgoDBName string
)

// Open 连接数据库
func Open(conf *Config) error {
	var (
		source string
		err    error
	)
	if conf.Type == "mysql" {
		source = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&allowOldPasswords=1&parseTime=true",
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
	if debug != nil {
		db.LogMode(*debug)
	}
	// 同步数据库
	if err = db.AutoMigrate(&Group{}, &Admin{}, &AdminLog{}, &Node{}).Error; err != nil {
		log.Printf("failed migrate (%s)", err.Error())
		return err
	}
	// 加载节点数据
	if mapNodes, err = loadNodes(); err != nil {
		log.Printf("failed init nodes (%s)", err.Error())
		return err
	}

	// 链接Redis
	if conf.RedisConf != nil {
		redisPool = redis.NewPool(conf.RedisConf)
	}

	if conf.MongoConf != nil {
		mgo = mongo.NewPool(conf.MongoConf)
		mgoDBName = conf.MongoConf.DBName
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
