package model

import (
	"encoding/hex"
	"encoding/json"
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
func (m Menu) Assign(g int64, n *Node) map[string]interface{} {
	return map[string]interface{}{"m": m, "Group": g, "Node": n}
}

func loadNodes() (map[int64]*Node, error) {
	var list []*Node
	n := map[int64]*Node{0: &Node{ID: 0, Path: "#"}}
	err := db.Order("id").Preload("Groups").Find(&list).Error
	if err != nil {
		return nil, err
	}
	for _, node := range list {
		n[node.ID] = node
	}
	for _, node := range list {
		if node.ID > 0 && node.Status {
			p := n[node.Parent]
			p.Child = append(p.Child, node)
		}
	}
	return n, nil
}

// Install 初始化节点
func Install(u *Admin, path string) error {
	data, err := ioutil.ReadFile("nodes.json")
	if err != nil {
		return err
	}
	if err = json.Unmarshal(data, &u.Group.Nodes); err != nil {
		return err
	}
	db := db.Begin().Set("gorm:association_autoupdate", true)
	u.Group.ID = 1
	if err = db.Save(&u.Group).Error; err != nil {
		db.Rollback()
		return err
	}
	u.ID = 1
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

// GetNodeAllNodes 根据用户组获取节点
func GetNodeAllNodes() (Menu, error) {
	var list Menu
	err := db.Order("id").Preload("Groups").Find(&list).Error
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

func init() {
	model = append(model, new(Node))
}
