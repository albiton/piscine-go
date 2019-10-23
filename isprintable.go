package piscine

func IsPrintable(str string) bool {
	if str == "" {
		return false
	}
	s := []rune(str)
	for _, letter := range s {
		if letter >= 'a' && letter <= 'z' {
			continue
		} else if letter >= 'A' && letter <= 'Z' {
			continue
		} else {
			return false
		}
	}
	return true
}
