package main
import (
	"fmt"
)

func main(){
	var total int = 0
	for a := 1; a < 1000; a++ {
		if a % 3 == 0 || a % 5 == 0 {
			total = total + a
		}
	}
	fmt.Printf("total: %d\n",total)
}
