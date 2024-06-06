package revision

func LastRune(s string) rune {
	last := []rune(s)

	return last[len(s)-1]
}
