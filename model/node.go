package model

import (
	"encoding/hex"
	"gocms/util"
	"io/ioutil"
	"os"

	"github.com/gorilla/securecookie"
	"gopkg.in/yaml.v2"
)

// NodeType 节点类型
type NodeType int8

const (
	// NodeTypeNormal 普通节点
	NodeTypeNormal NodeType = iota
	// NodeTypeEssensial 必要节点
	NodeTypeEssensial
	// NodeTypeAdmin 管理员
	NodeTypeAdmin
)

// Node 节点模型
type Node struct {
	ID     int      `gorm:"primary_key;auto_increment" yaml:"-"`
	Name   string   `gorm:"size:64;not null"`
	Parent int      `gorm:"default:0;not null" yaml:"-"`
	Icon   string   `gorm:"size:32;default:null" yaml:",omitempty"`
	Remark string   `gorm:"type:text" yaml:",omitempty"`
	Path   string   `gorm:"size:255;default:null"`
	Type   NodeType `gorm:"default:0;not null" yaml:",omitempty"`
	Status bool     `gorm:"default:false;not null"`
	Child  Menu     `yaml:",omitempty"`
	Groups []Group  `gorm:"many2many:node_groups" yaml:"-"`
}

// Menu 菜单
type Menu []*Node

// Assign 用于递归生成菜单
func (m Menu) Assign(g int64, n *Node) map[string]interface{} {
	return map[string]interface{}{"m": m, "Group": g, "Node": n}
}

func loadNodes() (map[int]*Node, error) {
	var list Menu
	n := map[int]*Node{0: &Node{ID: 0, Path: "#"}}
	err := db.Order("id").Preload("Groups").Find(&list).Error
	if err != nil {
		return nil, err
	}
	for _, node := range list {
		n[node.ID] = node
	}
	for _, node := range list {
		if node.ID > 0 {
			p := n[node.Parent]
			p.Child = append(p.Child, node)
		}
	}
	return n, nil
}

// Install 初始化节点
func Install(u *Admin, path string) error {
	data, err := ioutil.ReadFile("nodes.yml")
	if err != nil {
		return err
	}
	var (
		walk func(m Menu, id int) int
		list Menu
	)
	if err = yaml.Unmarshal(data, &list); err != nil {
		return err
	}
	walk = func(m Menu, id int) int {
		var i = id
		for _, v := range m {
			v.Parent = id
			i = i + 1
			if v.ID = i; len(v.Child) > 0 {
				i = walk(v.Child, v.ID)
				v.Child = nil
			}
			u.Group.Nodes = append(u.Group.Nodes, v)
		}
		return i
	}
	walk(list, 0)

	db := db.Begin().Set("gorm:association_autoupdate", true)
	if err = db.Save(&u.Group).Error; err != nil {
		db.Rollback()
		return err
	}
	u.GroupID = u.Group.ID
	u.Salt = hex.EncodeToString(securecookie.GenerateRandomKey(5))
	u.Password = util.MD5(u.Password + util.MD5(u.Salt))
	u.Status = true
	if err = db.Save(u).Error; err != nil {
		db.Rollback()
		return err
	}
	if err = db.Commit().Error; err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	if err = yaml.NewEncoder(f).Encode(&Config); err != nil {
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
func (n *Node) Parents() Menu {
	list := make(Menu, 0)
	for n.Parent != 0 {
		n = mapNodes[n.Parent]
		list = append(Menu{n}, list...)
	}
	return list
}

// HasGroup 判断指定节点是否能被某角色访问
func (n *Node) HasGroup(id int64) bool {
	if n == nil || n.Type == NodeTypeEssensial {
		return true
	}
	for _, role := range n.Groups {
		if role.ID == id {
			return true
		}
	}
	return false
}
