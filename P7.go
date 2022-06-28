package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

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
func NthPrime(N int) int {
	var sieve = Sieve(N)
	var limit =  N
	var len_primes = len(sieve)

	for len_primes < N{
		sieve = Sieve(limit)
		len_primes = len(sieve)
		limit *=2
	}
	return sieve[N-1]
}
func main(){
	var _n,_ = strconv.Atoi(os.Args[1])
	fmt.Printf("[*] %dst Prime number :> %d\n",_n,NthPrime(_n))
	return
}
