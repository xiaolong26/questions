package main

import "fmt"
//二叉树层次遍历
type queue struct {
	dequenode *Treenode
	next *queue
}

type Treenode struct {
	val int
	left *Treenode
	right *Treenode
}

func Offer(que *queue,node *Treenode)  {
	head := &queue{node,que}
	que = head
}

func Poll(que *queue) *Treenode{
	for  {
		que = que.next
		if que.next.next ==nil{
			break
		}
	}
	treenode := que.dequenode
	que.next = nil
	return treenode
}

func LevelTraverse(root *Treenode)  {
	que := &queue{root,nil}
	for  {
		deque := Poll(que)
		fmt.Println(deque.val)
		Offer(que,deque.left)
		Offer(que,deque.right)
		if que==nil{
			return
		}
	}
}

func main()  {
	root := &Treenode{1,nil,nil}
	LevelTraverse(root)
}