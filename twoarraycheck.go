package main

func seek(marr [][]int,target ...[]int)(x int,y int){
	if len(marr) == 0 || len(marr[0]) == 0{
		return -1,-1
	}
	y = len(marr[0])
	x = len(marr)
	return x,y
}

func main()  {

}
