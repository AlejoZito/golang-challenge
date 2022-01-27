package main

import "fmt"

var unorderedArray = [99]int{2, 3, 4, 5, 6, 7, 1, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
	46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70,
	71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95,
	21, 22, 23, 24, 25, 26, 27, 28 /*29,*/, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45,
	96, 97, 98, 99, 100}

//Challenge: provided there's an unordered list with 99 numbers, starting from 1 up to 100 and one number is missing:
//Find the missing number with a solution that prioritizes performance

func main() {
	//expected total if none was missing: 1+2+3+...+100
	expected := 0
	arraySum := 0

	for i := 0; i < 100; i++ {
		expected += i + 1

		if len(unorderedArray) > i {
			arraySum += unorderedArray[i]
		}
	}

	missingNumber := expected - arraySum

	fmt.Printf("the missing number is %d", missingNumber)
}
