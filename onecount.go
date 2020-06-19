package main

import "fmt"
//二进制几个1
func NumberOf1(n int)int{
	x := uint8(n)
	count := 0
	/*与c语言while(x)判断值得布尔类型不同，要判断值是否等于*/
	for x!=0{
		count += 1
		x = (x-1)&x
	}
	return count
}
func main(){
	count := NumberOf1(2)
	fmt.Println(count)
}
