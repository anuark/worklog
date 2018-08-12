// https://www.hackerrank.com/challenges/breaking-best-and-worst-records/problem
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var maxCap = 1000 * 1024

func main() {
	buf := make([]byte, maxCap)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCap)
	scanner.Scan()

	scanner.Scan()
	scores := strings.Split(scanner.Text(), " ")
	var lastScores []float64
	var breakMaxRecord int
	var breakMinRecord int
	var currentMaxScore float64
	var currentMinScore float64
	for i, v := range scores {
		// fmt.Println(i)
		v, _ := strconv.Atoi(v)
		score := float64(v)
		if i == 0 {
			currentMaxScore = score
			currentMinScore = score
		} else {
			if currentMaxScore < score {
				// fmt.Println(score, "breaked max record", currentMaxScore)
				breakMaxRecord++
			}

			if currentMinScore > score {
				// fmt.Println(score, "breaked min record", currentMinScore)
				breakMinRecord++
			}

			// fmt.Println("Max record", math.Max(score, currentMaxScore), "between", float64(score), currentMaxScore)
			currentMaxScore = math.Max(score, currentMaxScore)
			// fmt.Println("Min record", math.Min(score, currentMinScore), "between", float64(score), currentMinScore)
			currentMinScore = math.Min(score, currentMinScore)
		}

		lastScores = append(lastScores, score)
	}

	// fmt.Println(" ")
	fmt.Printf("%v %v", breakMaxRecord, breakMinRecord)
}
