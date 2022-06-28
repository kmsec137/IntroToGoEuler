package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func SumOfPrimes(limit int) int{
	var i int = 0
	var sieve []int = Sieve(limit)
	var total int = 0
	fmt.Println("[*] sieve:",sieve)
	for i = 0; i < len(sieve); i++{
		total = total + sieve[i]
	}
	return total
}

func Sieve(N int) []int {
	var sieve []int
	N = N + 1
	sieve = make([]int, N,N)
	var j int = 0
	var i int = 0
	var limit int  = N/2

	if N > 1000000 {
		limit = int(math.Sqrt(float64(N)))
	}

	sieve[0] = 1
	sieve[1] = 1
	sieve[2] = 0

	for i = 2; i < len(sieve);i++{
		sieve[i] = 0
	}
	for i = 2; i < limit; i++{
		if sieve[i] == 0 {
			for j = i*i;j < N;j+=i{
				sieve[j] = 1
			}
		}
	}
	PrintEro(sieve)
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
	fmt.Printf("[*] Sum Of Primes Below %d : %d\n",limit,SumOfPrimes(limit))
	return
}
