package main

import (
	"fmt"
	"strconv"
)

func isNumberic(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func getNumber(s string) float64 {
	var isNumber bool
	var n string
	for !isNumber {
		fmt.Printf("%s = ", s)
		fmt.Scan(&n)
		isNumber = isNumberic(n)
	}
	fn, _ := strconv.ParseFloat(n, 8)
	return fn
}

func getOperator() (string, error) {
	var op string
	var isOperator bool
	for !isOperator {
		fmt.Print("type operator +, -, * or / ")
		fmt.Scan(&op)
		isOperator = operatorChecker(op)
	}
	return op, nil
}

func operatorChecker(s string) bool {
	switch s {
	case
		"+",
		"-",
		"*",
		"/":
		return true
	}
	return false
}

func calc(x float64, y float64, op string) (float64, error) {
	var result float64
	switch op {
	case "+":
		result = x + y
	case "-":
		result = x - y
	case "*":
		result = x * y
	case "/":
		if y == 0 {
			result = 0
		} else {
			result = x / y
		}
	}
	return result, nil
}

func getPrimeNum() (int64, error) {
	var isPositive bool
	var s string
	for !isPositive {
		fmt.Printf("type positive number = ")
		fmt.Scan(&s)
		isPositive, _ = isPositiveNum(s)
		if !isPositive {
			fmt.Println("Error: You need to type positive number")
		}
	}
	var result, _ = strconv.ParseInt(s, 10, 64)
	return result, nil
}

func isPositiveNum(s string) (bool, error) {
	n, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		if n >= 0 {
			return true, nil
		}
	}
	return false, nil
}

func getListOfPrimeNumers(n int64) []int64 {
	var result []int64
	var x int64
	for x = 0; x <= n; x++ {
		var isPrime = checkIfPrime(x)
		if isPrime {
			result = append(result, x)
		}
	}
	return result
}

func checkIfPrime(n int64) bool {
	var x int64
	if n <= 1 {
		return false
	}
	for x = 2; x < n; x++ {
		if n%x == 0 {
			return false
		}
	}
	return true
}

func main() {
	// task 1
	fmt.Println("Task 1")
	fmt.Println("Calculator")

	var v1 = getNumber("x")
	var op, _ = getOperator()
	var v2 = getNumber("y")

	res, calcErr := calc(v1, v2, op)
	if calcErr != nil {
		fmt.Println("Error: check the input")
	} else {
		fmt.Printf("%g %s %g = %g \n", v1, op, v2, res)
	}
	fmt.Println("")

	// task 2
	fmt.Println("Task 2")
	fmt.Println("Get list of prime numbers from 0 till n")
	var primeNum, primeErr = getPrimeNum()
	if primeErr != nil {
		fmt.Println("Error: check the input")
	} else {
		var primes = getListOfPrimeNumers(primeNum)
		fmt.Printf("List of all prime numbers from 0 till %d \n", primeNum)
		fmt.Println(primes)
	}
}