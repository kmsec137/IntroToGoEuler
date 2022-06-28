package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	var index int = 0
	var limit int = 0
	limit,_ = strconv.Atoi(os.Args[1])
	var sum_of_squares int = 0
	var sum int = ((limit + 1)*limit)/2
	var square_of_sum int = sum*sum
	
	for index = 0; index <= limit; index++{
		sum_of_squares = index*index + sum_of_squares
	}
	fmt.Printf("[*] sum of squares = %d\n square of sum = %d\ndifference => %d\n",sum_of_squares,square_of_sum,-1*sum_of_squares + square_of_sum)
}
