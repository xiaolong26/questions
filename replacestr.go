package main

import "fmt"
//字符串替换
func replace(str *string)string{
	var r string
	for _,v :=range *str{
		if v==32{
			r+="%20"
		}else {
			r+=string(v)
		}
	}
	return string(r)
}

func main(){
	//长一点的要修改的字符串可以用[]rune
	str := "hello world"
	number := "123123"
	a := []rune(number)
	x := a[0] - '0'
	fmt.Println(x)
	fmt.Println(a[0]+a[1])
	r := replace(&str)
	fmt.Println(r)
}