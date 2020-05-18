package googp

// toSnake converts the string into a snake case.
func toSnake(str string) string {
	runes := []rune(str)
	var p int
	for i := 0; i < len(runes); i++ {
		c := runes[i]
		if c >= 'A' && c <= 'Z' {
			runes[i] = c - ('A' - 'a')
			if p+1 < i {
				tmp := append([]rune{'_'}, runes[i:]...)
				runes = append(runes[0:i], tmp...)
			}
			p = i
		}
	}
	return string(runes)
}

// isUpperPrefix returns true if the string starts with an uppercase letter.
// Otherwise, it returns false.
func isUpperPrefix(str string) bool {
	runes := []rune(str)
	if len(runes) == 0 {
		return false
	}

	c := runes[0]
	if c >= 'A' && c <= 'Z' {
		return true
	}
	return false
}
