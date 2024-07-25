package rsv

func LastWord(s string) string {
	check := false
	word := ""
	for _, char := range s {
		if char != ' ' {
			check = true
		}
	}
	if check {
		for i := len(s) - 1; i >= 0; i-- {
			if word == "" && s[i] == ' ' {
				continue
			}
			if word != "" && s[i] == ' ' {
				break
			}
			word = string(s[i]) + word
		}
	}
	return word + "\n"
}
