package search

func FindMissingNumberCheapSolution(arrayToSearch []int) int {
	//expected total if none was missing: 1+2+3+...+100
	expected := 0
	arraySum := 0

	for i := 0; i < len(arrayToSearch)+1; i++ {
		expected += i + 1

		if len(arrayToSearch) > i {
			arraySum += arrayToSearch[i]
		}
	}

	return expected - arraySum
}
