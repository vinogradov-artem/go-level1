package main

import (
	"fmt"
	"log"
	"math/big"
	"runtime"
	"time"
)

func FibonacciFun() func() *big.Int {
	var first, second *big.Int = big.NewInt(1), big.NewInt(1)
	return func() *big.Int {
		res := first
		first, second = second, new(big.Int).Add(first, second)
		return res
	}
}

func FibonacciDefered() func() *big.Int {
	var first, second *big.Int = big.NewInt(1), big.NewInt(1)
	return func() *big.Int {
		defer func() {
			first = new(big.Int).Add(first, second)
			first, second = second, first
		}()
		return first
	}
}

func FibonacciCached(n int, cache map[int]*big.Int) *big.Int {
	//	defer timeTrack(time.Now(), "")
	if value, found := cache[n]; found {
		return value
	}
	cache[n] = new(big.Int).Add(FibonacciCached(n-1, cache), FibonacciCached(n-2, cache))
	return cache[n]
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	if len(name) == 0 {
		name = funcName(2)
	}
	log.Printf("function %s took %v to complete.", name, elapsed)
}

func funcName(callerNum int) string {
	pc, _, _, _ := runtime.Caller(callerNum)
	return runtime.FuncForPC(pc).Name()
}

func main() {
	var n int
	for {
		fmt.Println("Enter number (int) for Fibonacci: ")
		if c, e := fmt.Scanf("%d", &n); e != nil && c != 1 {
			fmt.Printf("Invalid input %v - try again.\n", n)
			continue
		}
		break
	}

	{
		defer timeTrack(time.Now(), "FibonacciFun")
		f := FibonacciFun()
		for i := 1; i < n; i++ {
			f()
		}
		fmt.Println("FibonacciFun: ", f())
	}

	{
		defer timeTrack(time.Now(), "FibonacciDefered")
		f := FibonacciDefered()
		for i := 1; i < n; i++ {
			f()
		}
		fmt.Println("FibonacciDefered: ", f())
	}

	{
		defer timeTrack(time.Now(), "FibonacciCached")
		fmt.Println("FibonacciCached: ", FibonacciCached(n, map[int]*big.Int{1: big.NewInt(1), 2: big.NewInt(1)}))
	}

}
