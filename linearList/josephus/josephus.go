package josephus

import "fmt"

type Node struct {
	value int
	Next  *Node
}

var pHead *Node

func CreateJosephus(num int) {
	pHead = new(Node)
	var pS *Node
	var pEnd *Node
	pHead.value = 1
	for i := 2; i <= num; i++ {
		pS = new(Node)
		if i == 2 {
			pHead.Next = pS
			pS.value = i
			pEnd = pS
			continue
		}
		pS.value = i
		pEnd.Next = pS
		pEnd = pS
	}
	pEnd.Next = pHead
}

func Suicide(flag int, num int, p *Node) {
	for num != 0 {
		for i := 1; i < flag-1; i++ { //flag-1的目的是使指针停在目标前一位，避免指针指到目标而访问不到前驱，因为是单向链表
			p = p.Next
		}
		fmt.Println(p.Next.value)
		p.Next = p.Next.Next
		p = p.Next
		num--
	}
}
