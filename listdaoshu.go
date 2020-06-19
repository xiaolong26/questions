package main

import "fmt"
//删除链表倒数节点
type ListNode struct{
	next *ListNode
	key int
}

func FindKthToTail(head *ListNode,k int)int{
	a := &ListNode{}
	var b *ListNode
	c := []int{1,2,3,4,5,6}
	for i,j:=range c{
		fmt.Println(i,j)
	}
	a = head
	b = head
	for i := 0; i < k; i++ {
		a = a.next			
	}
	for a!=nil{
		b = b.next
		a = a.next
	}
	return b.key
}

func main(){
	arr := make(map[int]*ListNode)
	chal := make(chan int,10)
	chal <- 1
	select {
		case i := <-chal:
			fmt.Println(i)
	}
	arr[0] = &ListNode{nil,0}
	for i := 1; i < 10; i++ {
		arr[i] = &ListNode{arr[i-1],i}
	}
	result := FindKthToTail(arr[9],4)
	c := []int{1,2,3,4,5,6}
	for i,j:=range c{
		fmt.Println(i,j)
	}
	fmt.Println(result)
}