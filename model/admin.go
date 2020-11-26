package model

import (
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"gocms/util"
	"log"
	"strings"
	"time"

	"golang.org/x/xerrors"
	"gorm.io/gorm"
)

const ipChangeLimit = 3 // 信任IP数量

// AdminFlag 管理员选项
type AdminFlag int

const (
	_                     AdminFlag = iota
	FlagResetPassNext               = 1 // 下次必须重置密码
	FlagPassNeverExpire             = 2 // 密码永不过期
	FlagSecondLoginVerify           = 4 // 二次登入验证
)

// Admin 管理员
type Admin struct {
	ID        int64          `gorm:"primaryKey"`
	Email     string         `gorm:"size:255;uniqueIndex;not null"`
	Password  string         `gorm:"size:64;not null" json:"-"`
	Salt      string         `gorm:"size:10;not null" json:"-"`
	GroupID   int64          `gorm:"not null"`
	Headpic   string         `gorm:"size:255" json:",omitempty"`
	LastIP    string         `gorm:"size:16" json:",omitempty"`
	Flags     AdminFlag      `gorm:"default:1;not null"`
	LastLogin *time.Time     `json:",omitempty"`
	CreatedAt *time.Time     `gorm:"not null"`
	UpdatedAt *time.Time     `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
	Group     Group          `gorm:"-"`
}

func (m *Admin) String() string {
	return m.Email
}

// Status 判断管理员状态
func (m *Admin) Status() bool {
	if m.Flags&FlagResetPassNext > 0 {
		return false
	}
	return true
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
	r := db.Unscoped().Model(m).
		Where("email = ?", m.Email).
		Updates(map[string]interface{}{
			"group_id":   m.GroupID,
			"deleted_at": nil})
	if r.Error != nil {
		return r.Error
	}
	if r.RowsAffected > 0 {
		return nil
	}
	return db.Create(m).Error
}

// Delete 删除
func (m *Admin) Delete() error {
	return db.Delete(m).Error
}

// Update 更新
func (m *Admin) Update(v ...string) error {
	return db.Model(m).Select(v).Updates(m).Error
}

// Access 该用户能否访问指定节点
func (m *Admin) Access(tpl string) bool {
	if node := GetNodeByPath(tpl); node != nil {
		return node.HasGroup(m.GroupID)
	}
	log.Printf("Access `%s`", tpl)
	return true
}

// AdminLogin 登录请求
type AdminLogin struct {
	Email    string
	Password string
	IP       string
	UA       string
	Verifyed bool // 已通过短信验证
}

// Login 用户登录
func (p *AdminLogin) Login() (*Admin, error) {
	var m Admin
	if err := db.Take(&m, "email = ?", p.Email).Error; err != nil {
		return nil, errors.New("用户不存在")
	}
	if err := db.Take(&m.Group, m.GroupID).Error; err != nil {
		return nil, errors.New("用户组不存在")
	}
	if util.MD5(p.Password+util.MD5(m.Salt)) != m.Password {
		return nil, errors.New("密码不正确")
	}
	// 二次登录验证
	if !p.Verifyed && (m.Flags&FlagSecondLoginVerify > 0) {
		if err := JudgeAdmin(&m, p.IP); err != nil {
			return &m, xerrors.Errorf("JudgeAdmin %w", err)
		}
	}
	// 记录此次登录 IP
	err := db.Transaction(func(tx *gorm.DB) error {
		ctx := tx.Model(&m).UpdateColumns(map[string]interface{}{
			"last_ip": p.IP, "last_login": time.Now().UTC()})
		if ctx.Error != nil {
			return ctx.Error
		}
		return tx.Create(&AdminRecord{AdminID: m.ID, IP: p.IP,
			UA: p.UA, Action: actionLogin}).Error
	})
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
	ID    int64  `gorm:"primaryKey"`
	Name  string `gorm:"size:64;unique;not null"`
	Nodes Menu   `gorm:"many2many:node_groups"`
}

// Create 新建用户组
func (m *Group) Create() error {
	err := db.Create(m).Error
	if err != nil {
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
	err := db.Model(m).Association("Nodes").Replace(m.Nodes)
	if err != nil {
		return err
	}
	err = db.Model(m).Omit("Nodes").Updates(m).Error
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
	CreatedAt *time.Time `gorm:"not null" xlsx:"时间"`
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

// AdminAction 管理员行为
type AdminAction int

const (
	_ AdminAction = iota
	actionLogin
	actionResetPassward
	actionBindPhone
)

// AdminRecord 管理员操作记录
type AdminRecord struct {
	ID        int64       `gorm:"primary_key;auto_increment" xlsx:"-"`
	AdminID   int64       `gorm:"not null" xlsx:"-"`
	Action    AdminAction `gorm:"int4;index" xlsx:"动作"`
	IP        string      `gorm:"size:16" xlsx:"IP"`
	UA        string      `gorm:"size:255" xlsx:"-"`
	CreatedAt *time.Time  `gorm:"type:timestamp" xlsx:"时间"`
}

// JudgeAdmin 判断管理员登录IP
func JudgeAdmin(u *Admin, ip string) error {
	if u.LastIP == ip {
		return nil
	}
	// 判断近期登录 IP
	var recent []AdminRecord
	err := db.Model(&AdminRecord{}).Limit(ipChangeLimit).
		Select([]string{`ip`, `MAX("id") "id"`}).Group("ip").
		Order("id DESC").Find(&recent, "admin_id = ?", u.ID).Error
	if err != nil {
		return err
	}
	if len(recent) == 0 {
		return nil
	}
	for _, v := range recent {
		if v.IP == ip {
			return nil
		}
	}
	return errors.New("IPChange")
}

func init() {
	gob.Register(new(Admin))
}
