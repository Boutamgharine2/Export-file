package ascii


func Checkinput(s string) bool {
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		if r[i] < 32 || r[i] > 127 { // Hello
			if r[i] == 10 {
				return true
			} else {
				return false
			}
		}
	}
	return true
}
