package main

import "fmt"
//用两个栈模拟队列
type t []int
var stack1 t
var stack2 t
//方法接收器的赋值不会传播到其他调用,要用t得地址类型
func (stack *t)PopStack()(bool,int){
	if stack==nil{
		return false,0
	}
	x := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return true,x
}

func (stack *t) PushStack (x int){
	*stack = append(*stack,x)
}

func PopQueue()(bool,int){
	if len(stack2)!=0 {
		ok,val := stack2.PopStack()
		return ok,val
	}
	for {
		if len(stack1)==0{
			break
		}
		_,val := stack1.PopStack()
		stack2.PushStack(val)
	}
	ok,val := stack2.PopStack()
	return ok,val
}

func PushQueue(val int){
	stack1.PushStack(val)
}

func main(){
	for i:=0;i<10;i++{
		PushQueue(i)
	}
	for i:=0;i<5;i++{
		_,cal := PopQueue()
		fmt.Println(cal)
	}
	fmt.Println(len(stack1))
}