package piscine

func StrRev(s string) string {
	var rev string
	for _, k := range s {
		rev = string(k) + rev
	}
	return rev
}
