package main

import (
	"fmt"
)

func Pic(dx,dy int) [][]uint8  {
	pow := make([][]uint8,uint8(dy))
	for i:=range pow {
		pow[i] = make([]uint8,uint8(dx))
		for j :=range pow[i]{
			pow[i][j] = uint8(dx)
		}
	}
	return pow
}

func main() {
	p := []int{2,3,5,7,11,13}
	fmt.Println("p ==",p)
	fmt.Println("p[1:4]==",p[1:4])
	fmt.Println("p[:3]==",p[:3])
	fmt.Println("p[4:]==",p[4:])
	fmt.Println(Pic(3,2))

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
func fibonacci() func() int  {
	a,b := 0,1
	return func() int{
		a,b = b,a+b
		return a
	}
}
