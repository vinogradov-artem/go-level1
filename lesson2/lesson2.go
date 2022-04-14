package main

import (
	"fmt"
	"math"
	"strconv"
)

func getRectangleSides() (x, y int64) {
	var w, l int64
	fmt.Println("Lets count area of the rectangle")
	fmt.Print("Enter width of the rectangle: ")
	fmt.Scan(&w)
	fmt.Print("Enter lenght of the rectangle: ")
	fmt.Scan(&l)
	return w, l
}

func RectangleArea(x, y int64) (int64, error) {
	return x * y, nil
}

func RectanglePerimetar(x, y int64) (int64, error) {
	return 2 * (x + y), nil
}

func getRadiusOfCircleArea() (float64, error) {
	var a float64
	fmt.Print("Enter area of the circle: ")
	fmt.Scanln(&a)
	var as string = fmt.Sprint(a)
	fmt.Println("Area of the circle is: " + string(as))
	var r float64 = math.Sqrt(a / math.Pi)
	return r, nil
}

// диаметр круга
func CircleDiametar(r float64) (float64, error) {
	return 2 * r, nil
}

// длину окружности
func CirclePerimetar(r float64) (float64, error) {
	return 2 * float64(r) * math.Pi, nil
}

func hundredsTensUnits(num int64) (string, error) {
	fmt.Printf("Your number is: %d \n", num)

	var result string

	if num >= 100 && num <= 999 {
		mystring := strconv.FormatInt(int64(num), 10)

		for i := 0; i < len(mystring); i++ {

			switch i {
			case 0:
				result += "hundreds = " + string(mystring[i]) + "\n"
			case 1:
				result += "tens = " + string(mystring[i]) + "\n"
			case 2:
				result += "units = " + string(mystring[i])
			}
		}
	}
	return result, nil
}

func getNumber() {
	var mynum int64
	var properNumber bool = false
	for !properNumber {
		fmt.Print("Enter number between 100 and 999: ")
		fmt.Scanln(&mynum)
		properNumber, _ = checkNumber(mynum)
	}
	htu, _ := hundredsTensUnits(mynum)
	fmt.Println(htu)
	fmt.Println("")
}

func checkNumber(x int64) (bool, error) {
	var result bool = false
	if x >= 100 && x <= 999 {
		result = true
	}
	return result, nil
}

func main() {
	// task 1
	fmt.Println("Task 1")
	var m, n = getRectangleSides()
	fmt.Println("")

	fmt.Printf("Rectangle: widht = %d, lenght = %d \n", m, n)

	var rectanglePerimetar, _ = RectanglePerimetar(m, n)
	fmt.Println("Rectangle Perimetar = ", rectanglePerimetar)

	var rectangleArea, _ = RectangleArea(m, n)
	fmt.Println("Rectangle Area = ", rectangleArea)

	fmt.Println("")

	// task 2
	fmt.Println("Task 2")
	var r, _ = getRadiusOfCircleArea()
	var d, _ = CircleDiametar(r)
	var c, _ = CirclePerimetar(r)

	fmt.Println("Radius of the circle - r = ", r)
	fmt.Println("Diametar of the circle - d = ", d)
	fmt.Println("Perimetar of the circle - C = ", c)
	fmt.Println("")

	// task 3
	fmt.Println("task 3")
	getNumber()
}
