package main

import "fmt"

type ListNode3 struct {
	Next *ListNode3
	val int
}

func main()  {
	firstnode := &ListNode3{nil,1}
	dict := make(map[int]*ListNode3)
	dict[1] = firstnode
	for i:=2;i<=10;i++{
		dict[i] = &ListNode3{dict[i-1],i}
	}
	firstnode = dict[10]
	/*head := DeleteListnode(&firstnode,dict[1])
	for {
		fmt.Println(head)
		head = head.Next
		if head==nil{
			break
		}
	}*/
	lastnode := &ListNode3{nil,100}
	fmt.Println(DeleteListnode(&lastnode,lastnode))
}

func DeleteListnode(head **ListNode3,deletenode *ListNode3)*ListNode3{
	innode := *head
	headback := *head
	if *head==nil || deletenode==nil{
		return nil
	}
	if deletenode ==headback{
		innode = nil
	}else if deletenode.Next == nil {
		for {
			if headback.Next != deletenode {
				headback = headback.Next
			} else {
				break
			}
		}
	}else {
		deletenode.val = deletenode.Next.val
		deletenode.Next = deletenode.Next.Next
	}
	return innode
}