// https://www.hackerrank.com/challenges/grading/problem
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	for scanner.Scan() {
		inGrade := scanner.Text()
		grade, _ := strconv.Atoi(inGrade)
		fmt.Println(roundGrade(grade))
		// break
	}
}

func roundGrade(grade int) int {
	if grade < 38 {
		return grade
	}

	for i := 40; i <= 100; i += 5 {
		// fmt.Println(grade, i > grade, math.Abs(float64(grade)-float64(i)) < 3)
		if i > grade && math.Abs(float64(grade)-float64(i)) < 3 {
			return i
		}
	}

	return grade
}
