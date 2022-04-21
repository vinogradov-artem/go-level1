package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var numbers string
	var listNumbers []int

	fmt.Println("Введите целые числа через запятую:")
	scannerNumber := bufio.NewScanner(os.Stdin)
	scannerNumber.Scan()
	numbers = scannerNumber.Text()

	arrayNumbers := strings.Split(numbers, ",")
	for _, v := range arrayNumbers {
		newValue := strings.TrimSpace(v)
		v, err := strconv.Atoi(newValue)
		if err != nil {
			fmt.Printf("Не является целым числом, ошибка: %s\n", err)
			os.Exit(1)
		}
		listNumbers = append(listNumbers, v)
	}

	sortNumbers(&listNumbers)
	fmt.Printf("результат %v", listNumbers)
}

func sortNumbers(listNumbers *[]int) {
	listNumbersSort := *listNumbers
	for i := 1; i < len(listNumbersSort); i++ {
		k := i
		for k > 0 && listNumbersSort[k-1] > listNumbersSort[k] {
			tmp := listNumbersSort[k-1]
			listNumbersSort[k-1] = listNumbersSort[k]
			listNumbersSort[k] = tmp
			k--
		}
	}
}
