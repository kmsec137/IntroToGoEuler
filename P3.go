package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func GreatestPrimeFactor(C int) int{
	var sqrt_limit int = int(math.Sqrt(float64(C)))
	var i int = 0
	var sieve []int = Sieve(sqrt_limit)
	var max int = sieve[0]
	fmt.Println("[*] sieve:",sieve)
	for i = 0; i < len(sieve); i++{
		//fmt.Printf("[i] [c: %d] mod [s: %d] :=> %d\n",C,sieve[i],C%sieve[i])
		if C%sieve[i] == 0 {
			if  sieve[i] > max {
				max = sieve[i]
			}
		}
	}
	return max
}
func Sieve(N int) []int {
	var sieve []int
	sieve = make([]int, N,N)
	var j int = 0
	var i int = 0
	var limit int  = 0

	sieve[0] = 1
	sieve[1] = 1
	sieve[2] = 0
	for i = 2; i < len(sieve);i++{
		sieve[i] = 0
	}
	limit = int(math.Sqrt(float64(N)))
	for i = 2; i < limit; i++{
		if sieve[i] == 0{
			for j = i*i;j < N;j+=i{
				sieve[j] = 1
			}
		}
	}
	var _sieve []int
	for i = 2;i < len(sieve); i++ {
		if sieve[i] == 0{
			_sieve = append(_sieve,i) 
		}
	}
	return _sieve
}
func PrintEro(sieve []int){
	var i int
	for i = 1; i < len(sieve);i++{
		if sieve[i] == 0{
			fmt.Printf("%.2d ",i)
		}
	}
	fmt.Println()
}

func usage(){

	fmt.Printf("[x] Usage: %s [limit]",os.Args[0])
}
func main(){
	var limit int
	var _ error

	if len(os.Args) != 2{
		usage()
	}
	fmt.Printf("[*] seive of erostothenes for limit: %s\n",os.Args[1])
	limit,_ = strconv.Atoi(os.Args[1])
	fmt.Printf("[*] greatest prime factor of %d : %d\n",limit,GreatestPrimeFactor(limit))
	return
}
