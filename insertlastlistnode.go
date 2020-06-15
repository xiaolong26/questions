package main

import "fmt"

type ListNode1 struct {
	next *ListNode1
	val int
}
//如果头节点为空，则head只能在函数内改变，所有要用指针的指针类型
//画个图就很清楚
func InsertlistnodeInlast(head **ListNode1)*ListNode1{
	last := &ListNode1{nil,100}
	//head为指针的指针，*head为该节点的地址，**head为该节点实际的值
	fu := *head
	if *head==nil {
		*head = last
		return *head
	}
	for {
		if fu.next==nil{
			fu.next = last
			break
		}
		fu = fu.next
	}
	return *head
}

func main()  {
	dict := make(map[int]*ListNode1)
	firstnode := &ListNode1{nil,10}
	dict[10] = firstnode
	for i:=9;i>0;i--{
		dict[i] = &ListNode1{dict[i+1],i}
	}
	/*for i:=1;i<10;i++ {
		fmt.Println(dict[i],dict[i].next)
	}*/
	a := dict[1]
	InsertlistnodeInlast(&a)
	fmt.Println(dict[10])
}
