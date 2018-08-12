package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	scanner.Scan()
	inp1 := strings.Split(scanner.Text(), " ")
	var aList []int
	for i := 0; i < len(inp1); i++ {
		v, _ := strconv.Atoi(inp1[i])
		aList = append(aList, v)
	}

	scanner.Scan()
	inp2 := strings.Split(scanner.Text(), " ")
	var bList []int
	for i := 0; i < len(inp2); i++ {
		v, _ := strconv.Atoi(inp2[i])
		bList = append(bList, v)
	}

	var betweenAB int
	for k := 1; k <= 100; k++ {
		var aDivCount int
		var bDivCount int

		for i := 0; i < len(aList); i++ {
			if k%aList[i] == 0 {
				aDivCount++
			}
		}

		for i := 0; i < len(bList); i++ {
			if bList[i]%k == 0 {
				bDivCount++
			}
		}

		// fmt.Println(k, aDivCount == len(aList)-1, bDivCount == len(bList)-1)
		if aDivCount == len(aList) && bDivCount == len(bList) {
			betweenAB++
		}
	}

	fmt.Println(betweenAB)
}
