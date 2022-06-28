package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	var cur int = 1
	var next int = 1
	var temp int = 0
	var limit,_ = strconv.Atoi(os.Args[1])
	var total = 0

	for next < limit {

		temp = cur + next
		cur = next
		next = temp
		fmt.Printf("[i] (%d, %d)\n",cur,next)
		if next%2 == 0 {
			total = total + next
		}
	}
	fmt.Printf("[*] total := %d\n",total)
}
