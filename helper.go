package helper

import "fmt"

func CalculateSum(x int, y int) int {
	var sum int
	sum = x + y
	return sum
}

func GreetUsers(conferenceName string) {
	fmt.Printf("Welcome to %v\n", conferenceName)
}
