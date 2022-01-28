package search

import "errors"

// FindMissingWithSort It's not technically a sort because it never obtains a sorted array,
// but for this use case an array with a false value in the missing number position is sufficient
func FindMissingWithSort(arrayToSearch []int) (int, error) {
	foundValues := findMissingPosition(arrayToSearch)

	for i := 0; i < len(foundValues); i++ {
		if !foundValues[i] {
			return i + 1, nil
		}
	}

	return 0, errors.New("something unexpected happened")
}

/*
Create a new array which will have a "true" value in every position where a corresponding value is found.
The missing number will be represented by a "false" value in the (missingNumber-1) position
*/
func findMissingPosition(arrayToSort []int) (foundValues []bool) {
	foundValues = make([]bool, len(arrayToSort)+1)

	for i := 0; i < len(arrayToSort); i++ {
		valueFound := arrayToSort[i]
		foundValues[valueFound-1] = true //inserting the value found instead of "true" would give us a sorted array
	}

	return foundValues
}
