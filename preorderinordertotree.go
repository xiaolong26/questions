package main

import "fmt"
//前序中序遍历确定树
type TreeNode struct {
	left *TreeNode
	right *TreeNode
	val int
}

func MakeTree(preord []int,inorder []int,root *TreeNode)*TreeNode{
	//root := &TreeNode{nil,nil,preord[0]}
	if len(preord)==0{
		return nil
	}
	if len(preord) == 1 {
		node := new(TreeNode)
		node.val = preord[0]
		return node
	}
	var target int
	for i,n := range preord{
		if n==preord[0]{
			target = i
		}
	}
	root.left = &TreeNode{nil,nil,0}
	root.right = &TreeNode{nil,nil,0}
	if target==len(preord)-1{
		root.right = nil
	}else if target==0 {
		root.left = nil
	}
	MakeTree(preord[1:target+1],inorder[0:target],root.left)
	MakeTree(preord[target+1:len(preord)],inorder[target+1:len(inorder)-1],root.right)
	return root
}

func main()  {
	precord := []int{1,2,4,7,3,5,6,8}
	inorder := []int{4,7,2,1,5,3,8,6}
	root := &TreeNode{val: 1}
	MakeTree(precord,inorder,root)
	fmt.Println(root.right)
	fmt.Println(root.left)
}