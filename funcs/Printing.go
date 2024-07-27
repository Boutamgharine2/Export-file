package ascii

import (
	"fmt"
	"os"
	"strings"
)

func Printing(s, r string) string {
	s = strings.ReplaceAll(s, "\r\n", "\n")
	if Checkinput(s) {
		f, err := os.Open("Fonts/" +r+".txt")
		if err != nil {
			fmt.Println("error :", err)
		}
		

		

		
		v := strings.Split(s, "\n")
		h := PrintN(f, v)
		return h
	}
	return "you need to choose a character from the ascii table."
}
