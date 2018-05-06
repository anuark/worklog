package main

import (
	"fmt"
	"strings"
)

func wordWrap(input string, size int) (text string, rows int) {
	sep := strings.Split(input, " ")

	var strLen int
	rows = 1
	for i, v := range sep {

		strLen += len(v) * wordSize
		if strLen > size && i != 0 {
			sep[i-1] = sep[i-1] + "\n"
			strLen = len(v) * wordSize
			rows++
		}

		sep[i] += " "
	}

	text = strings.Join(sep, "")
	return
}

func main() {
	str := "Murcielago venenoso maton peligroso mata caballo patas de cabro"
	fmt.Println(wordWrap(str, 22))
}
