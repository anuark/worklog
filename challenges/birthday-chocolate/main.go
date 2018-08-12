// https://www.hackerrank.com/challenges/the-birthday-bar/problem
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
	chocNumInp := strings.Split(scanner.Text(), " ")
	chocNum := []int{}
	for _, v := range chocNumInp {
		n, _ := strconv.Atoi(v)
		chocNum = append(chocNum, n)
	}

	scanner.Scan()
	mdInp := strings.Split(scanner.Text(), " ")
	m, _ := strconv.Atoi(mdInp[0])
	_ = m
	d, _ := strconv.Atoi(mdInp[1])
	var waysToBreakChoc int
	// fmt.Println(chocNum, m)
	for i := 0; i < len(chocNum); i++ {
		if len(chocNum) == 1 {
			if chocNum[i] == m {
				waysToBreakChoc++
			}
			break
		}

		if i < d-1 {
			continue
		}

		// fmt.Println(i, sum(chocNum[i-(d-1):i+1]...), chocNum[i-(d-1):i+1], i-(d-1), i+1, sum(chocNum[i-(d-1):i+1]...) == m)
		if sum(chocNum[i-(d-1):i+1]...) == m {
			waysToBreakChoc++
		}
	}

	fmt.Println(waysToBreakChoc)
}

func sum(nums ...int) int {
	var total int
	for _, v := range nums {
		total += v
	}
	return total
}
