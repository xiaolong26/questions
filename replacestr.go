package main

import "fmt"

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
	str := "hello world"
	r := replace(&str)
	fmt.Println(r)
}