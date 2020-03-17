package model

import (
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"gocms/util"
	"log"
	"strings"
	"time"

	"github.com/gorilla/securecookie"
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
	LastLogin *time.Time
	CreatedAt *time.Time `gorm:"not null"`
	UpdatedAt *time.Time `json:"-"`
	DeletedAt *time.Time `json:"-"`
	Group     Group      `gorm:"-"`
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
	m.Headpic = "/static/img/avatar.png"
	db := db.Unscoped().Model(m).
		Where("email = ?", m.Email).
		Updates(map[string]interface{}{
			"group_id":   m.GroupID,
			"deleted_at": nil})
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected > 0 {
		return nil
	}
	return db.Create(m).Error
}

// Delete 删除
func (m *Admin) Delete() error {
	return db.Delete(m).Error
}

// UpdatePasswd 更新密码
func (m *Admin) UpdatePasswd(v ...interface{}) error {
	m.Salt = hex.EncodeToString(securecookie.GenerateRandomKey(5))
	m.Password = util.MD5(m.Password + util.MD5(m.Salt))
	v = append(v, "salt")
	if !m.Status {
		m.Status = true
		v = append(v, "status")
	}
	return db.Model(m).Select("password", v...).Updates(m).Error
}

// Access 该用户能否访问指定节点
func (m *Admin) Access(tpl string) bool {
	if node := GetNodeByPath(tpl); node != nil {
		return node.HasGroup(m.GroupID)
	}
	log.Printf("Access `%s`", tpl)
	return true
}

// Login 用户登录
func Login(email, passwd, ip string) (*Admin, error) {
	var m Admin
	if err := db.Take(&m, "email = ?", strings.ToLower(email)).Error; err != nil {
		return nil, errors.New("用户不存在")
	}
	if err := db.Take(&m.Group, m.GroupID).Error; err != nil {
		return nil, errors.New("用户组不存在")
	}
	if m.Status && util.MD5(passwd+util.MD5(m.Salt)) != m.Password {
		return nil, errors.New("密码不正确")
	}
	err := db.Model(&m).UpdateColumns(map[string]interface{}{
		"last_ip": ip, "last_login": time.Now()}).Error
	return &m, err
}

// GetAdmins 获取用户列表
func GetAdmins(filter ...func(*gorm.DB) *gorm.DB) ([]*Admin, error) {
	var list []*Admin
	err := db.Scopes(filter...).Order("id").Find(&list).Error
	return list, err
}

// GetAdminNum 获取用户数量
func GetAdminNum(filter ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var nums int64
	err := db.Model(&Admin{}).Scopes(filter...).Count(&nums).Error
	return nums, err
}

// Group 用户组
type Group struct {
	ID    int64  `gorm:"primary_key;auto_increment"`
	Name  string `gorm:"size:64;unique;not null"`
	Nodes Menu   `gorm:"many2many:node_groups;association_autoupdate:false"`
}

// Create 新建用户组
func (m *Group) Create() error {
	err := db.Select("id").Find(&m.Nodes, "type = ?", NodeTypeEssensial).Error
	if err != nil {
		return err
	}
	if err = db.Create(m).Error; err != nil {
		return err
	}
	if mapNodes, err = loadNodes(); err != nil {
		return fmt.Errorf("init nodes failed: %v", err)
	}
	return nil
}

func (m *Group) String() string {
	return m.Name
}

// Select 获取角色
func (m *Group) Select() error {
	return db.Take(m).Error
}

// Update 更新角色
func (m *Group) Update() error {
	err := db.Model(m).Association("Nodes").Replace(m.Nodes).Error
	if err != nil {
		return err
	}
	err = db.Model(m).Set("gorm:association_save_reference", false).Updates(m).Error
	if err != nil {
		return err
	}
	if mapNodes, err = loadNodes(); err != nil {
		return err
	}
	return nil
}

// GetGroups 获取角色列表
func GetGroups() (map[int64]string, error) {
	var (
		list []*Group
		hash = make(map[int64]string)
	)
	if err := db.Order("id").Find(&list).Error; err != nil {
		return nil, err
	}
	for _, v := range list {
		hash[v.ID] = v.Name
	}
	return hash, nil
}

// AdminLog 操作日志
type AdminLog struct {
	ID        int64      `gorm:"primary_key;auto_increment" xlsx:"-"`
	AdminID   int64      `gorm:"not null" xlsx:"-"`
	Admin     *Admin     `xlsx:"用户"`
	Path      string     `gorm:"size:255;not null" xlsx:"路径"`
	UA        string     `gorm:"size:255" xlsx:"-"`
	Commit    string     `gorm:"type:text" xlsx:"注释"`
	IP        string     `gorm:"size:16" xlsx:"IP"`
	CreatedAt *time.Time `gorm:"type(datetime)" xlsx:"时间"`
}

// Create 插入日志
func (m *AdminLog) Create() error {
	return db.Create(m).Error
}

// GetLogs 获取日志列表
func GetLogs(filter ...func(*gorm.DB) *gorm.DB) ([]AdminLog, error) {
	var list []AdminLog
	err := db.Scopes(filter...).Preload("Admin", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, email")
	}).Order("id desc").Find(&list).Error
	return list, err
}

// GetLogNum 获取日志数量
func GetLogNum(filter ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var nums int64
	err := db.Model(&AdminLog{}).Scopes(filter...).Count(&nums).Error
	return nums, err
}

func init() {
	gob.Register(new(Admin))
}
