package redis

import (
	"backend/domain"
	"testing"
	"time"
)

func Test_NewPool(t *testing.T) {
	conf := domain.DBConf{
		Host:   "localhost",
		Port:   6379,
		DBName: "0",
	}
	r := NewPool(&conf)

	key := "UNITTEST"
	expire := time.Second * 10
	conn := r.NewConn()
	if err := conn.Set(key, 1, expire).Err(); err != nil {
		t.Error(err)
	}
	if err := conn.Get(key).Err(); err != nil {
		t.Error(err)
	}
}
