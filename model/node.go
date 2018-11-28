package model

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"github.com/dragonflylee/gocms/util"
)

// NodeType 节点类型
type NodeType int8

const (
	// NodeTypeNormal 普通节点
	NodeTypeNormal NodeType = iota
	// NodeTypeEssensial 必要节点
	NodeTypeEssensial
	// NodeTypeFix 不可修改节点
	NodeTypeFix
)

// Node 节点模型
type Node struct {
	ID     int64    `gorm:"primary_key;auto_increment"`
	Name   string   `gorm:"size:64;not null"`
	Parent int64    `gorm:"default:0;not null"`
	Icon   string   `gorm:"size:32;default:null"`
	Remark string   `gorm:"type:text"`
	Path   string   `gorm:"size:255"`
	Type   NodeType `gorm:"default:0;not null"`
	Status bool     `gorm:"default:false;not null"`
	Child  Menu     `gorm:"-"`
	Groups []*Group `gorm:"many2many:node_groups"`
}

// Menu 菜单
type Menu []*Node

// Assign 用于递归生成菜单
func (m Menu) Assign(group int64, node *Node) map[string]interface{} {
	return map[string]interface{}{"m": m, "group": group, "node": node}
}

func loadNodes() (map[int64]*Node, error) {
	var (
		list []*Node
		m    = map[int64]*Node{0: &Node{ID: 0}}
	)
	err := db.New().Order("id").Preload("Groups").Find(&list).Error
	if err != nil {
		return nil, err
	}
	for _, node := range list {
		m[node.ID] = node
	}
	for _, node := range list {
		if node.ID > 0 && node.Status {
			p := m[node.Parent]
			p.Child = append(p.Child, node)
		}
	}
	return m, nil
}

// Install 初始化节点
func Install(m *Admin, path string) error {
	data, err := ioutil.ReadFile(filepath.Join(path, "nodes.json"))
	if err != nil {
		return err
	}
	if err = json.Unmarshal(data, &m.Group.Nodes); err != nil {
		return err
	}
	db := db.New().Begin().Set("gorm:association_autoupdate", true)
	if err = db.Create(&m.Group).Error; err != nil {
		db.Rollback()
		return err
	}
	m.GroupID = m.Group.ID
	m.Salt = util.RandString(10)
	m.Password = util.Md5Hash(m.Password + util.Md5Hash(m.Salt))
	m.Status = true
	if err = db.Create(m).Error; err != nil {
		db.Rollback()
		return err
	}
	if err = db.Commit().Error; err != nil {
		return err
	}
	if mapNodes, err = loadNodes(); err != nil {
		return err
	}
	return nil
}

// GetNodes 获取节点树
func GetNodes() Menu {
	return mapNodes[0].Child
}

// GetNodeAllNodes 根据用户组获取节点
func GetNodeAllNodes() (Menu, error) {
	var list Menu
	err := db.New().Order("id").Preload("Groups").Find(&list).Error
	if err != nil {
		return nil, err
	}
	dict := map[int64]*Node{0: &Node{ID: 0}}
	for _, node := range list {
		dict[node.ID] = node
	}
	for _, node := range list {
		p := dict[node.Parent]
		p.Child = append(p.Child, node)
	}
	return dict[0].Child, nil
}

// GetNodeByPath 根据路径查找节点
func GetNodeByPath(path string) *Node {
	for _, node := range mapNodes {
		if node.Path == path {
			return node
		}
	}
	return nil
}

// HasParent 判断父节点是否存在
func (n *Node) HasParent(id int64) bool {
	if n == nil {
		return false
	}
	for n.ID != 0 {
		if n.ID == id {
			return true
		}
		n = mapNodes[n.Parent]
	}
	return false
}

func (n *Node) String() string {
	if n == nil {
		return ""
	}
	return n.Name
}

// Parents 获取指定节点的所有父节点
func (n *Node) Parents() []*Node {
	list := make([]*Node, 0)
	for n.Parent != 0 {
		n = mapNodes[n.Parent]
		list = append([]*Node{n}, list...)
	}
	return list
}

// HasGroup 判断指定节点是否能被某角色访问
func (n *Node) HasGroup(id int64) bool {
	if n == nil {
		return true
	}
	for _, role := range n.Groups {
		if role.ID == id {
			return true
		}
	}
	return false
}

// Group 用户组
type Group struct {
	ID    int64   `gorm:"primary_key;auto_increment"`
	Name  string  `gorm:"size:64;unique_index;not null"`
	Nodes []*Node `gorm:"many2many:node_groups;association_autoupdate:false"`
}

// Create 新建用户组
func (m *Group) Create() error {
	err := db.New().Select("id").Find(&m.Nodes, "type = ?", NodeTypeEssensial).Error
	if err != nil {
		return err
	}
	return db.New().Create(m).Error
}

func (m *Group) String() string {
	return m.Name
}

// Select 获取角色
func (m *Group) Select() error {
	return db.New().Take(m).Error
}

// Update 更新角色
func (m *Group) Update() error {
	err := db.New().Model(m).Association("Nodes").Replace(m.Nodes).Error
	if err != nil {
		return err
	}
	err = db.New().Model(m).Set("gorm:association_save_reference", false).Updates(m).Error
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
	if err := db.New().Order("id").Find(&list).Error; err != nil {
		return nil, err
	}
	for _, v := range list {
		hash[v.ID] = v.Name
	}
	return hash, nil
}
