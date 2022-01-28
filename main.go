package main

import (
	"fmt"
	"golang-challenge/file"
	"golang-challenge/search"
)

//Challenge: provided there's an unordered list with 99 numbers, starting from 1 up to 100 and one number is missing:
//Find the missing number with a solution that prioritizes performance
func main() {
	unorderedArray, err := file.ReadJson("unorderedArray.json")
	if err != nil {
		fmt.Errorf("failed reading json file")
	}

	cheapSolutionMissingNumber := search.FindMissingNumberCheapSolution(unorderedArray)
	fmt.Printf("using the cheap solution, missing number is %d \n", cheapSolutionMissingNumber)

	sortSolutionMissingNumber, _ := search.FindMissingWithSort(unorderedArray)
	fmt.Printf("using the sort solution, missing number is %d \n", sortSolutionMissingNumber)

}
