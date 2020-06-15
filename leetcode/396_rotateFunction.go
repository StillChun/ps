func maxRotateFunction(A []int) int {
	if len(A) <= 1 {
		return 0
	}

	sum := 0
	tmp := 0
	answer := -0xFFFFFFFF

	for i, v := range A {
		sum += v
		tmp += i * v
	}

	for _, v := range A {
		tmp = tmp - sum + len(A)*v
		if tmp > answer {
			answer = tmp
		}
	}
	return answer
}