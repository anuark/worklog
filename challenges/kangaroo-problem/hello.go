// https://www.hackerrank.com/challenges/kangaroo/problem
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const maxCap = 100 * 1024

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, maxCap)
	scanner.Buffer(buf, maxCap)
	scanner.Scan()
	vals := scanner.Text()
	valsSep := strings.Split(vals, " ")
	start1, _ := strconv.Atoi(valsSep[0])
	jump1, _ := strconv.Atoi(valsSep[1])
	start2, _ := strconv.Atoi(valsSep[2])
	jump2, _ := strconv.Atoi(valsSep[3])

	kangaroo1Pos := start1
	kangaroo2Pos := start2
	for i := 0; i < 10000; i++ {
		kangaroo1Pos += jump1
		kangaroo2Pos += jump2
		if kangaroo1Pos == kangaroo2Pos {
			fmt.Println("YES")
			return
		}
	}

	fmt.Println("NO")
}
