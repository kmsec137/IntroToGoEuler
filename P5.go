package main

import (
	"fmt"
	"os"
	"strconv"
)

func factorial(N int64) int64 {
	var i int64 = 0
	var _f int64 = 1
	for i = N; i > 1;i--{
		_f  = _f*i
	}
	return _f
}
func isDivisible(C int64) int{
	var i int64
	for i = 20; i > 1; i++{
		if C%i != 0{
			return 0
		}
	}
	return 1
}

func main(){
	var index int64 = 0
	var f, _= strconv.Atoi(os.Args[1])
	var factorial_limit = factorial(int64(f))

	fmt.Printf("[*] factorial of (%d) :=> %d\n",f,factorial_limit)
	var smallest int64 = factorial_limit
	for index = factorial_limit; index > int64(f); index--{
		if isDivisible(index) == 1{
			fmt.Println("[*] found one: %d",index)
			if index < smallest{
				smallest = index
			}
		}
	}
}
