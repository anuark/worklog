package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Murcielago venenoso maton peligroso mata caballo patas de cabro"
	a := wordWrap(str, 22)
	// fmt.Printf("%#v", a)
	fmt.Println(a)
}

var wordSize = 1

func wordWrap(input string, size int) string {
	sep := strings.Split(input, " ")

	var strLen int
	addSpace := true
	for i, v := range sep {
		if addSpace {
			sep[i] += " "
		}

		addSpace = true

		strLen += len(v) * wordSize
		if strLen > size && i != 0 {
			sep[i-1] = sep[i-1] + "\n"
			strLen = len(v) * wordSize
			addSpace = false
		}
	}

	return strings.Join(sep, "")
}
