package damlev

// DamerauLevenshteinDistance calculates the edit distances between two strings,
// counting the required number of insertions, deletions, mutations, or transpositions.
func DamerauLevenshteinDistance(a, b string) int {
	aLen := len(a)
	bLen := len(b)
	if aLen == 0 {
		return bLen
	} else if bLen == 0 {
		return aLen
	}
	_a, _b := " "+a, " "+b
	aLen, bLen = len(_a), len(_b)

	distanceMatrix := make([][]int, bLen)
	for i := 0; i < bLen; i++ {
		distanceMatrix[i] = make([]int, aLen)
	}

	for i := 0; i < aLen; i++ {
		distanceMatrix[0][i] = i
	}
	for i := 0; i < bLen; i++ {
		distanceMatrix[i][0] = i
	}
	for j := 1; j < bLen; j++ {
		for i := 1; i < aLen; i++ {
			if _a[i] == _b[j] {
				distanceMatrix[j][i] = distanceMatrix[j-1][i-1]
			} else {
				distanceMatrix[j][i] = 1 + min(distanceMatrix[j][i-1], distanceMatrix[j-1][i], distanceMatrix[j-1][i-1])
			}

			if _a[i] == _b[j-1] && _a[i-1] == _b[j] {
				distanceMatrix[j][i] = min(distanceMatrix[j][i], distanceMatrix[j-2][i-2]+1)
			}
		}
	}
	return distanceMatrix[bLen-1][aLen-1]
}

func min(first int, ints ...int) int {
	minimum := first
	for _, n := range ints {
		if n < minimum {
			minimum = n
		}
	}
	return minimum
}
