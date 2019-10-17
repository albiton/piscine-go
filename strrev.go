package piscine

func StrRev(s string) string {
	var reverse string
	for i, j := 0, len(s)-1; i > 0; i-- {
		reverse += string(s[i])
	}
	return reverse
}
