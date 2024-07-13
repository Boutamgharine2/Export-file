package ascii

func Checker(s []string) bool {
	for _, t := range s {
		if t != "" {
			return false
		}
	}
	return true
}
