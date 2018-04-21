package model

import (
	"crypto/md5"
	"encoding/gob"
	"encoding/hex"
	"errors"
	"time"
)

// Admin 管理员
type Admin struct {
	ID        int64  `gorm:"primary_key;auto_increment"`
	Email     string `gorm:"size:255;unique_index;not null"`
	Password  string `gorm:"size:64;not null"`
	Salt      string `gorm:"size:10;not null"`
	GroupID   int64  `gorm:"not null"`
	Headpic   string `gorm:"size:255"`
	LastIP    string `gorm:"size:16"`
	LastLogin time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	Group     Group
}

// Create 注册新用户
func (m *Admin) Create() error {
	m.Password = encryptPass(m.Password, m.Salt)
	m.Headpic = "/static/img/avatar.png"
	return db.New().Create(m).Error
}

// Login 用户登录
func Login(email, passwd, ip string) (*Admin, error) {
	var (
		user Admin
		db   = db.New()
	)
	if db.Where("email = ?", email).First(&user).RecordNotFound() {
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
	update := map[string]interface{}{
		"last_ip": ip, "last_login": time.Now(),
	}
	if err := db.Model(&user).Updates(update).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// AdminLog 操作日志
type AdminLog struct {
	ID        int64  `gorm:"primary_key;auto_increment"`
	AdminID   int64  `gorm:"not null"`
	Path      string `gorm:"size:255;not null"`
	Commit    string `gorm:"type:text"`
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

// encryptPass 生成密码串
func encryptPass(password, salt string) string {
	return md5Hash(password + md5Hash(salt))
}

func init() {
	gob.Register(&Admin{})
}
