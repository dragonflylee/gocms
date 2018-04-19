package model

import (
	"crypto/md5"
	"encoding/gob"
	"encoding/hex"
	"errors"
	"math/rand"
	"time"
)

// Admin 管理员
type Admin struct {
	ID         int64  `gorm:"primary_key,auto_increment"`
	Username   string `gorm:"size:64"`
	Email      string `gorm:"size:255"`
	Password   string `gorm:"size:64"`
	Repassword string `gorm:"-"`
	Salt       string `gorm:"size:10"`
	GroupID    int64
	Headpic    string `gorm:"size:255"`
	LastIP     string `gorm:"size:16"`
	LastLogin  time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Group      Group
}

// Create 注册新用户
func (m *Admin) Create() error {
	m.Salt = randString(10)
	m.Password = encryptPass(m.Password, m.Salt)
	m.Headpic = "/static/img/avatar.png"
	return db.New().Create(m).Error
}

// Login 用户登录
func Login(name, passwd, ip string) (*Admin, error) {
	var (
		user Admin
		db   = db.New()
	)
	if db.Where("username = ?", name).First(&user).RecordNotFound() {
		return nil, errors.New("用户不存在")
	}
	if db.Model(&user).Related(&user.Group).RecordNotFound() {
		return nil, errors.New("用户组不存在")
	}
	if db.Error != nil {
		return nil, db.Error
	}
	if encryptPass(passwd, user.Salt) != user.Password {
		return nil, errors.New("密码不正确")
	}
	db.Model(&user).Updates(map[string]interface{}{
		"last_ip": ip, "last_login": time.Now()})
	return &user, nil
}

// AdminLog 操作日志
type AdminLog struct {
	ID        int64 `gorm:"primary_key,auto_increment"`
	AdminID   int64
	Path      string `gorm:"size:255"`
	Commit    string
	IP        string `gorm:"size:16"`
	CreatedAt time.Time
}

// Create 插入日志
func (m *AdminLog) Create() error {
	return db.New().Create(m).Error
}

// md5Hash 生成32位MD5
func md5Hash(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

// randString 生成随机字符串
func randString(l int) string {
	bytes := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// encryptPass 生成密码串
func encryptPass(password, salt string) string {
	return md5Hash(password + md5Hash(salt))
}

func init() {
	gob.Register(&Admin{})
}
