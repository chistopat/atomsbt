package main

import (
	"fmt"
	"math/rand"
	"time"
)

const kMaxNumbersCount = 100

func Accumulate(array *[]int) int {
	acc := 0
	for i := 0; i < len(*array); i++ {
		acc += (*array)[i]
	}
	return acc
}

func FoundLostNumber(numbers *[]int) int {
	actualSum := Accumulate(numbers)

	// arithmetic progression
	targetSum := (len(*numbers) + 1) * len(*numbers) / 2
	return targetSum - actualSum
}

func GetRandomNumber(max int) int {
	rand.Seed(time.Now().UnixNano())
	left, right := 1, max+1
	return rand.Intn(right - left)
}

func main() {
	// make numbers from 1 to 100
	numbers := make([]int, kMaxNumbersCount)
	for i := 1; i <= kMaxNumbersCount; i++ {
		numbers[i-1] = i
	}

	fmt.Println(numbers)
	// erase random index by zero
	randomIndex := GetRandomNumber(kMaxNumbersCount)
	numbers[randomIndex] = 0
	fmt.Println(numbers)

	// random shuffle slice
	for i := range numbers {
		j := rand.Intn(i + 1)
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	fmt.Println(numbers)

	// found answer
	fmt.Println(FoundLostNumber(&numbers))
}
