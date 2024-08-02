package ascii

import (
	"fmt"
	"os"
	"strings"
)

func Printing(s, r string) string {
	s = strings.ReplaceAll(s, "\r\n", "\n")

	f, err := os.Open("Fonts/" + r + ".txt")
	if err != nil {
		fmt.Println("error :", err)
	}

	v := strings.Split(s, "\n")
	h := PrintN(f, v)
	return h
}
