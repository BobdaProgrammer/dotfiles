package main

import (
	"fmt"
	"math"
)

func main(){
	fmt.Println("Hello World")
	var arr []int = []int{0,2,3}
	fmt.Println(0)
	for i := range arr{
		if i!=0{
			fmt.Println(math.Abs(float64(arr[i]-arr[i-1])))		
		}
	}
}
