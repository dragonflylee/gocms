package model

import (
	_ "embed"

	"gocms/pkg/config"
	"gocms/pkg/util"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
)

// NodeType 节点类型
type NodeType int8

const (
	// NodeNormal 普通节点
	NodeNormal NodeType = iota
	// NodeAll 必要节点
	NodeAll
	// NodeAdmin 管理员
	NodeAdmin
)

// Node 节点模型
type Node struct {
	ID     int      `gorm:"primaryKey" yaml:"-"`
	Name   string   `gorm:"size:64;not null"`
	Parent int      `gorm:"default:0;not null" yaml:"-"`
	Icon   string   `gorm:"size:32;default:fa fa-circlo-o" yaml:",omitempty"`
	Remark string   `gorm:"type:text" yaml:",omitempty"`
	Path   string   `gorm:"size:255"`
	Type   NodeType `gorm:"default:0;not null" yaml:",omitempty"`
	Status bool     `gorm:"default:false;not null"`
	Child  Menu     `gorm:"-" yaml:",omitempty"`
	Groups []Group  `gorm:"many2many:node_groups" yaml:"-"`
}

// Menu 菜单
type Menu []*Node

func (m Menu) Recurve(data gin.H) gin.H {
	data["Child"] = m
	return data
}

func loadNodes() (map[int]*Node, error) {
	var menus Menu
	n := map[int]*Node{0: {ID: 0, Path: "#"}}
	err := db.Order("id").Preload("Groups").Find(&menus).Error
	if err != nil {
		return nil, err
	}
	for _, node := range menus {
		n[node.ID] = node
	}
	for _, node := range menus {
		if node.ID > 0 {
			p := n[node.Parent]
			p.Child = append(p.Child, node)
		}
	}
	return n, nil
}

func walkNode(root *Menu, m Menu, id int) int {
	var i = id
	for _, v := range m {
		v.Parent = id
		i = i + 1
		if v.ID = i; len(v.Child) > 0 {
			i = walkNode(root, v.Child, v.ID)
			v.Child = nil
		}
		*root = append(*root, v)
	}
	return i
}

//go:embed menu.yml
var menu []byte

// Install 初始化节点
func Install(u *Admin, path string) error {
	var menus Menu
	err := yaml.Unmarshal(menu, &menus)
	if err != nil {
		return err
	}

	walkNode(&u.Group.Nodes, menus, 0)

	err = db.Transaction(func(tx *gorm.DB) error {
		err := tx.Save(&u.Group).Error
		if err != nil {
			return err
		}
		u.GroupID = u.Group.ID
		u.Salt = util.RandString(5)
		u.Password = util.MD5(u.Password, util.MD5(u.Salt))
		u.Flags = FlagPassNeverExpire
		if err = tx.Save(u).Error; err != nil {
			return err
		}
		return config.Save(path)
	})
	if err != nil {
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
func (n *Node) HasParent(id int) bool {
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
func (n *Node) Parents() []string {
	var slice []string
	for n.Parent != 0 {
		n = mapNodes[n.Parent]
		slice = append(slice, n.Name)
	}
	return slice
}

// HasGroup 判断指定节点是否能被某角色访问
func (n *Node) HasGroup(id int64) bool {
	if n == nil || n.Type == NodeAll {
		return true
	}
	for _, role := range n.Groups {
		if role.ID == id {
			return true
		}
	}
	return false
}
