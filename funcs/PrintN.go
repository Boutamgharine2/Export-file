package ascii

import (
	"os"
)

func PrintN(inputfile *os.File, v []string) string {
	c := ""
	if Checker(v) {
		for k := 0; k < len(v)-1; k++ {
			c += "\n"
		}
		return c
	}
	for _, word := range v {
		var char []string
		var chars [][]string
		r := []rune(word)
		if word == "" {
			c += "\n" // hello\nworld
			continue
		}
		for i := 0; i < len(r); i++ {
			char = Printingchar(r[i], inputfile)
			chars = append(chars, char)
		}

		for l := 0; l < 8; l++ {
			for j := 0; j < len(chars); j++ {
				c += chars[j][l]
			}
			c += "\n"
		}
	}

	defer inputfile.Close()
	return c
}
