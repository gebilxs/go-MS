package goweb

import (
	"strings"
)

type treeNode struct {
	name     string
	children []*treeNode
}

//put path: /user/get/:id

func (t *treeNode) Put(path string) {
	//恢复现场
	root := t
	//对斜杠进行相对应的分割
	strs := strings.Split(path, "/")
	for index, name := range strs {
		if index == 0 {
			continue
		}
		children := t.children
		isMatch := false
		for _, node := range children {
			if node.name == name {
				t = node
				isMatch = true
				break
			}
		}
		if !isMatch {
			node := &treeNode{name: name, children: make([]*treeNode, 0)}
			children = append(children, node)
			t.children = children
			t = node
		}
	}
	t = root
}

//get path: /user/get/1

func (t *treeNode) Get(path string) *treeNode {
	strs := strings.Split(path, "/")
	for index, name := range strs {
		if index == 0 {
			continue
		}
		children := t.children
		isMatch := false
		for _, node := range children {
			if node.name == name || node.name == "*" || strings.Contains(node.name, ":") {
				t = node
				isMatch = true
				if index == len(strs)-1 {
					return node
				}
				break
			}
		}
		if !isMatch {
			for _, node := range children {
				if node.name == "**" {
					return node
				}
			}
		}
	}
	return nil
}
