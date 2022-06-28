package main
import (
	"fmt"
	"math"
	"strconv"
)

func main(){
	var power uint64 = uint64(math.Pow(2,1000))
	var number_string =  strconv.FormatUint(power,10)
	fmt.Printf("[*] 2**1000 = %s\n",number_string)
	var index int = 0
	var total int = 0
	var num = 0
	for index = 0;index < len(number_string); index ++{
		num,_ = strconv.Atoi(string(number_string[index]))
		total = total + int(num)
	}
	fmt.Printf("[*] total := %d\n",total)
	return
}
