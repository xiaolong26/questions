package main

import "fmt"
//按圈打印二维数组
func PrintArr(arr [][]int,colums int,rows int)  {
	start := 0
	//找出规律，行跟列大于start即可打印一圈，每次少一行跟一列
	for colums > start*2 && rows >start*2{
		//只传入start数值就可以打印一圈
		PrintCircle(arr,colums,rows,start)
		start += 1
	}
}

func PrintCircle(arr [][]int,colums int,rows int,start int){
	endx := colums-1-start
	endy := rows-1-start
	for i:=0;i<endx;i++{
		number := arr[start][i]
		fmt.Println(number)
	}
	for i:=0;i<endy;i++{
		number := arr[i][endx]
		fmt.Println(number)
	}
}

func main()  {

}