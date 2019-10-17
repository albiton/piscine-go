package piscine

<<<<<<< HEAD
func StrRev(s string) string {
=======
func StrRev (s string) string {
>>>>>>> 48ea4de6a0e622e5c693c1dd28ce2b5848038185
	var reverse string
	for i, j := 0, len(s)-1; i > 0; i-- {
		reverse += string(s[i])
	}
	return reverse
}
