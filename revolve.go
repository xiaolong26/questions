package main

import "fmt"

type ListNode2 struct {
	next *ListNode2
	val int
}

func main(){
	dict := make(map[int]*ListNode2)
	firstnode := &ListNode2{nil,0}
	dict[0] = firstnode
	for i:=1;i<=10;i++ {
		dict[i] = &ListNode2{dict[i-1],i}
	}
	RevolveList(dict[10])
}


//递归可以类比为栈，将函数放进栈中
func RevolveList(head *ListNode2){
	if head.next != nil{
		//压栈操作，这个函数后面语句要等递归完后函数才可以运行，先进后出
		RevolveList(head.next)
	}
	fmt.Println(head.val)
}