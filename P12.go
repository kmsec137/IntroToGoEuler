package main

import (
	"fmt"
	"os"
	"strconv"
)

func getDivisors(c int) []int {
	var index = 0
	var divisors = make([]int,c,c)
	for index = 1; index <= c;index++ {
		if (c%index == 0){
			divisors = append(divisors,index)
		}
	}
	return divisors
}

func listTriangleDivisors(limit int){
	var index int = 0
	var div int = 0
	var divisors []int
	for index = 1; index < limit; index ++{
		var tri = ((index + 1)*index)/2
		divisors = getDivisors(tri)
		fmt.Printf("%d: ",tri)
		for div = 0;div < len(divisors); div ++ {
			if divisors[div] != 0{
				fmt.Printf(" %d",divisors[div])
			}
		}
		fmt.Println()
	}
}

func fivehundredTriangleDivisors(){
	var index int = 1
	var div int = 0
	var divisors []int
	for index > 0{
		var tri = ((index + 1)*index)/2
		divisors = getDivisors(tri)
		if len(divisors) >= 500 {
			fmt.Printf("[*] %d: ",tri)
			for div = 0;div < len(divisors); div ++ {
				if divisors[div] != 0{
					fmt.Printf(" %d",divisors[div])
				}
			}
			fmt.Println()
			return
		}
	}
}

func main(){
	var limit int = 0
	limit, _ = strconv.Atoi(os.Args[1])
	if (limit > 1000) { fmt.Println() }
	fivehundredTriangleDivisors()
	return
}
