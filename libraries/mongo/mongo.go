package mongo

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"gopkg.in/mgo.v2"
)

type MongodbConfig struct {
	Host     string
	Port     int
	DBName   string
	Debug    bool
	Username string
	Password string
}

var GlobalSession *mgo.Session

const RetryCount = 5

func NewPool(conf *MongodbConfig) *mgo.Session {
	if conf.Host == "" || conf.Port <= 0 || conf.DBName == "" {
		panic(errors.New("MongoDB config error"))
	}
	mgoUrl := fmt.Sprintf("%s:%d?maxPoolSize=128", conf.Host, conf.Port)
	var (
		session *mgo.Session
		err     error
	)
	if conf.Debug {
		mgo.SetDebug(true)
		mLogger := log.New(os.Stderr, "", log.LstdFlags)
		mgo.SetLogger(mLogger)
	}
	for i := 0; i < RetryCount; i++ {
		mgo.SetDebug(conf.Debug)
		session, err = mgo.Dial(mgoUrl)
		if err != nil {
			log.Printf("Failed to connect mongodb: %v", conf)
			time.Sleep(2 * time.Second)
			log.Printf("Retrying to connect to mongodb: %v", conf)
			continue
		}
		session.SetPoolLimit(128)

		if conf.Username != "" && conf.Password != "" {
			err = session.DB(conf.DBName).Login(conf.Username, conf.Password)
			if err != nil {
				panic(err)
			}
		}

		return session
	}
	panic(err)
}

func NotFound(err error) bool {
	return err == mgo.ErrNotFound
}
