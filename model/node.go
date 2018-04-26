package model

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

// Node 节点模型
type Node struct {
	ID     int64    `gorm:"primary_key;auto_increment"`
	Name   string   `gorm:"size:64;not null"`
	Parent int64    `gorm:"default:0;not null"`
	Icon   string   `gorm:"size:16;default:null"`
	Path   string   `gorm:"size:255"`
	Remark string   `gorm:"type:text"`
	Status bool     `gorm:"default:false;not null"`
	Child  []*Node  `gorm:"-"`
	Groups []*Group `gorm:"many2many:node_groups"`
}

var (
	mapNodes map[int64]*Node
)

func loadNodes() error {
	var list []*Node
	if err := db.New().Order("id").Find(&list).Error; err != nil {
		return err
	}
	mapNodes = map[int64]*Node{0: &Node{ID: 0}}
	for _, node := range list {
		err := db.Model(node).Association("Groups").Find(&node.Groups).Error
		if err != nil {
			return err
		}
		mapNodes[node.ID] = node
	}
	for _, node := range list {
		if node.ID > 0 && node.Status {
			p := mapNodes[node.Parent]
			p.Child = append(p.Child, node)
		}
	}
	return nil
}

// Install 初始化节点
func Install(path string) error {
	var (
		list []*Node
		data []byte
		err  error
	)
	if data, err = ioutil.ReadFile(filepath.Join(path, "nodes.json")); err != nil {
		return err
	}
	if err = json.Unmarshal(data, &list); err != nil {
		return err
	}
	db := db.New().Begin()
	role := &Group{Name: "超级管理员"}
	if err = db.Create(role).Error; err != nil {
		db.Rollback()
		return err
	}
	for _, n := range list {
		n.Groups = append(n.Groups, role)
		if err = db.Create(n).Error; err != nil {
			db.Rollback()
			return err
		}
	}

	if err = db.Commit().Error; err != nil {
		return err
	}
	return loadNodes()
}

// GetNodes 获取节点树
func GetNodes() []*Node {
	return mapNodes[0].Child
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

// Parents 获取指定节点的所有父节点
func (n *Node) Parents() []*Node {
	list := make([]*Node, 0)
	for n.Parent != 0 {
		n = mapNodes[n.Parent]
		list = append(list, n)
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
	ID   int64  `gorm:"primary_key;auto_increment"`
	Name string `gorm:"size:64;not null"`
}

// Create 新建用户组
func (m *Group) Create() error {
	return db.New().Create(m).Error
}

func (m *Group) String() string {
	return m.Name
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
