package main

import (
	"fmt"
)

type BinTreeNode struct {
	ch                    string
	leftchild, rightchild *BinTreeNode
}

var flag = 0

//约定先序输入创建二叉树,Go中不存在加&传入引用的做法，递归的过程需要将结果return给孩子，实现初始化。否则最终根节点的孩子始终为nil
func CreateBinTree(p *BinTreeNode) *BinTreeNode {
	var c string
	fmt.Scan(&c)

	if c == "#" {
		p = nil
	} else {
		p = new(BinTreeNode)
		p.ch = c
		p.leftchild = CreateBinTree(p.leftchild)
		p.rightchild = CreateBinTree(p.rightchild)
	}
	return p
}

func PreOrder(p *BinTreeNode) {
	if p != nil {
		fmt.Printf("%s", p.ch)
		PreOrder(p.leftchild)
		PreOrder(p.rightchild)
	}
}

func InOrder(p *BinTreeNode) {
	if p != nil {
		InOrder(p.leftchild)
		fmt.Printf("%s", p.ch)
		InOrder(p.rightchild)
	}
}

func PostOrder(p *BinTreeNode) {
	if p != nil {
		PostOrder(p.leftchild)
		PostOrder(p.rightchild)
	    fmt.Printf("%s", p.ch)
	}
}

func main() {
	var root *BinTreeNode
	root = CreateBinTree(root)
	PreOrder(root)
	fmt.Println()
	InOrder(root)
	fmt.Println()
	PostOrder(root)
	fmt.Println()
}
