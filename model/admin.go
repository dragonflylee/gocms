package model

import (
	"crypto/md5"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"errors"
	"math/rand"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

// Admin 管理员
type Admin struct {
	ID        int64  `gorm:"primary_key;auto_increment"`
	Email     string `gorm:"size:255;unique_index;not null"`
	Password  string `gorm:"size:64;not null" json:"-"`
	Salt      string `gorm:"size:10;not null" json:"-"`
	GroupID   int64  `gorm:"not null"`
	Headpic   string `gorm:"size:255"`
	LastIP    string `gorm:"size:16"`
	Status    bool   `gorm:"default:false;not null"`
	LastLogin time.Time
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Group     Group     `gorm:"-"`
}

func (m *Admin) String() string {
	return m.Email
}

// GobEncode 序列化
func (m *Admin) GobEncode() ([]byte, error) {
	return json.Marshal(m)
}

// GobDecode 反序列化
func (m *Admin) GobDecode(data []byte) error {
	return json.Unmarshal(data, m)
}

// Create 注册新用户
func (m *Admin) Create() error {
	m.Email = strings.ToLower(m.Email)
	m.Salt = randString(10)
	m.Password = encryptPass(m.Password, m.Salt)
	m.Headpic = "/static/img/avatar.png"
	return db.New().Create(m).Error
}

// UpdatePasswd 更新密码
func (m *Admin) UpdatePasswd(v ...interface{}) error {
	m.Salt = randString(10)
	m.Password = encryptPass(m.Password, m.Salt)
	v = append(v, "salt")
	if !m.Status {
		m.Status = true
		v = append(v, "status")
	}
	return db.New().Model(m).Select("password", v...).Updates(m).Error
}

// Access 该用户能否访问指定节点
func (m *Admin) Access(tpl string) bool {
	if node := GetNodeByPath(tpl); node != nil {
		return node.HasGroup(m.GroupID)
	}
	return true
}

// Login 用户登录
func Login(email, passwd, ip string) (*Admin, error) {
	var (
		db   = db.New()
		user Admin
		err  error
	)
	if err = db.First(&user, "email = ?", email).Error; err != nil {
		return nil, errors.New("用户不存在")
	}
	if err = db.Model(&user).Related(&user.Group).Error; err != nil {
		return nil, errors.New("用户组不存在")
	}
	if encryptPass(passwd, user.Salt) != user.Password {
		return nil, errors.New("密码不正确")
	}
	err = db.Model(&user).UpdateColumns(&Admin{
		LastIP: ip, LastLogin: time.Now(),
	}).Error
	return &user, err
}

// GetAdmins 获取用户列表
func GetAdmins(filter ...func(*gorm.DB) *gorm.DB) (list []*Admin, err error) {
	err = db.New().Scopes(filter...).Order("id").Find(&list).Error
	return list, err
}

// GetAdminNum 获取用户数量
func GetAdminNum(filter ...func(*gorm.DB) *gorm.DB) (nums int64, err error) {
	err = db.New().Model(&Admin{}).Scopes(filter...).Count(&nums).Error
	return nums, err
}

// AdminLog 操作日志
type AdminLog struct {
	ID        int64     `gorm:"primary_key;auto_increment" xlsx:"-"`
	AdminID   int64     `gorm:"not null" xlsx:"-"`
	Admin     *Admin    `xlsx:"用户;field:Email"`
	Path      string    `gorm:"size:255;not null" xlsx:"路径"`
	UA        string    `gorm:"size:255" xlsx:"-"`
	Commit    string    `gorm:"type:text" xlsx:"注释"`
	IP        string    `gorm:"size:16" xlsx:"IP"`
	CreatedAt time.Time `xlsx:"时间"`
}

// Create 插入日志
func (m *AdminLog) Create() error {
	return db.New().Create(m).Error
}

// GetLogs 获取日志列表
func GetLogs(filter ...func(*gorm.DB) *gorm.DB) (list []*AdminLog, err error) {
	err = db.Scopes(filter...).Preload("Admin", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, email")
	}).Order("id desc").Find(&list).Error
	return list, err
}

// GetLogNum 获取日志数量
func GetLogNum(filter ...func(*gorm.DB) *gorm.DB) (nums int64, err error) {
	err = db.New().Model(&AdminLog{}).Scopes(filter...).Count(&nums).Error
	return nums, err
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

func init() {
	gob.Register(&Admin{})
}
