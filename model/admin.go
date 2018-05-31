package model

import (
	"encoding/gob"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"gocms/util"
)

// Admin 管理员
type Admin struct {
	ID        int64     `gorm:"primary_key;auto_increment"`
	Email     string    `gorm:"size:255;unique_index;not null"`
	Password  string    `gorm:"size:64;not null" json:"-"`
	Salt      string    `gorm:"size:10;not null" json:"-"`
	GroupID   int64     `gorm:"not null"`
	Headpic   string    `gorm:"size:255"`
	LastIP    string    `gorm:"size:16"`
	Status    bool      `gorm:"default:false;not null"`
	LastLogin time.Time `gorm:"type(datetime)"`
	CreatedAt time.Time `gorm:"type(datetime)" json:"-"`
	UpdatedAt time.Time `gorm:"type(datetime)" json:"-"`
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
	m.Salt = util.RandString(10)
	m.Password = util.Md5Hash(m.Password + util.Md5Hash(m.Salt))
	m.Headpic = "/static/img/avatar.png"
	m.LastLogin = time.Now()
	return db.New().Create(m).Error
}

// UpdatePasswd 更新密码
func (m *Admin) UpdatePasswd(v ...interface{}) error {
	m.Salt = util.RandString(10)
	m.Password = util.Md5Hash(m.Password + util.Md5Hash(m.Salt))
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
		db = db.New()
		m  = new(Admin)
	)
	if email = strings.TrimSpace(email); len(email) <= 0 {
		return nil, errors.New("邮箱格式不合法")
	}
	if db = db.Take(m, "email = ?", strings.ToLower(email)); db.Error != nil {
		return nil, errors.New("用户不存在")
	}
	if err := db.Related(&m.Group).Error; err != nil {
		return nil, errors.New("用户组不存在")
	}
	if util.Md5Hash(passwd+util.Md5Hash(m.Salt)) != m.Password {
		return nil, errors.New("密码不正确")
	}
	err := db.UpdateColumns(&Admin{LastIP: ip, LastLogin: time.Now()}).Error
	return m, err
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
	CreatedAt time.Time `gorm:"type(datetime)" xlsx:"时间"`
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

func init() {
	gob.Register(&Admin{})
}
