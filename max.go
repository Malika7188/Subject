package revision

func Max(a []int) int {
	var max int
	for _, char := range a {
		if char > max {
			max = char
		}
	}
	return max
}
