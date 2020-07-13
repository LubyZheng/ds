package linkList

import "fmt"

type Node struct {
	value int
	point *Node
}

var pHead *Node

func Create() *Node {
	pHead = new(Node)
	var val int
	var pS, pEnd *Node
	for {
		fmt.Scan(&val)
		if val == 0 {
			pEnd.point = nil
			return pHead
		}
		if pHead.point == nil {
			pS = new(Node)
			pS.value = val
			pHead.point = pS
			pEnd = pS
		} else {
			pS = new(Node)
			pS.value = val
			pEnd.point = pS
			pEnd = pS
		}
	}
}

func ShowList(p *Node) {
	for {
		if p.point == nil {
			break
		}
		fmt.Printf("%d ", p.point.value)
		p = p.point
	}
	fmt.Println()
}

func Add(p *Node, n int, val int) {
	for i := 0; i < n-1; i++ {
		p = p.point
		if p == nil && i != n-2 {
			fmt.Println("添加位置超出数组范围,添加失败")
			return
		}
	}
	pS := new(Node)
	pS.value = val
	pS.point = p.point
	p.point = pS
	return
}

func Delete(p *Node, n int) {
	for i := 0; i < n-1; i++ {
		p = p.point
		if p == nil && i != n-2 {
			fmt.Println("添加位置超出数组范围,添加失败")
			return
		}
	}
	p.point = p.point.point
	return
}
