import "sort"

func fairCandySwap(A []int, B []int) []int {
	var answer []int

	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})

	sort.Slice(B, func(i, j int) bool {
		return B[i] < B[j]
	})

	aSum := 0
	bSum := 0

	for _, v := range A {
		aSum += v
	}

	for _, v := range B {
		bSum += v
	}

	diff := aSum - bSum

	for i, j := 0, 0; i < len(A) && j < len(B); {
		nowDiff := A[i] - B[j]

		if diff == 2*nowDiff {
			answer = append(answer, A[i], B[j])
			break
		} else if diff > 2*nowDiff {
			i++
		} else {
			j++
		}
	}

	return answer
}