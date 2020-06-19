package main

import (
	"fmt"
)
//判断压栈顺序是否可以对应出栈顺序
func main()  {
	for i:=0;i<10;i++{
		Push(i)
	}
	for {
		i,ok := Pop()
		if !ok {
			break
		}
		fmt.Println(i)
	}
	pPush := []int{1,2,3}
	pPop := []int{3,2,1}
	StackOrder(pPush,pPop)
}

func StackOrder(pPush []int,pPop []int)bool{
	var count = 0
	res := false
	for _,val1 := range pPop{
		if count==0{
			for y,val2 := range pPush{
				if val2!=val1{
					Push(val2)
				}else {
					count = y
					break
				}
			}
		}else {
			for y:=count;y<len(pPush);y++{
				if pPush[y]==val1{
					Push(pPush[y])
					count = y
					res = true
					break
				}
				res = false
			}
			if !res{
				x,_:= Pop()
				if x!=val1 {
					return false
				}
			}
		}
	}
	return true
}

var stackback []int

func Push(val int){
	stackback = append(stackback,val)
}

func Pop()(int,bool){
	if len(stackback) <= 0{
		return 0,false
	}else{
		val := stackback[len(stackback)-1]
		stackback = stackback[:len(stackback)-1]
		return val,true
	}
}

