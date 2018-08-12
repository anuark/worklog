// https://www.hackerrank.com/challenges/mini-max-sum/problem
package main

import (
    "math"
    "fmt"
    "log"
    "bufio"
    "os"
    "strings"
    "strconv"
)

var _ = math.Pi
var _ = bufio.MaxScanTokenSize
var _ = os.DevNull
var _ = strings.Compare
var _ = strconv.IntSize
var _ = log.LUTC


func main() {
    scanner := bufio.NewScanner(os.Stdin)
    var (
        positiveCount int
        negativeCount int
        zeroesCount int
        size int
    )

    for scanner.Scan() {
        splitStr := strings.Split(scanner.Text(), " ")
        if len(splitStr) == 1 {
            // var s int
            s, err := strconv.Atoi(splitStr[0])
            if err != nil {
                log.Fatal(err)
            }
            
            size = s
            continue
        }

        for _, v := range splitStr {
            num, err := strconv.Atoi(v)
            if err != nil {
                log.Fatal(err)
            }
            // fmt.Println(num, num > 0)

            switch {
            case num > 0:
                positiveCount++
            case num < 0:
                negativeCount++
            default:
                zeroesCount++
            }
        }
    }

    // fmt.Println(size, positiveCount)
    fmt.Println(float64(positiveCount)/float64(size))
    fmt.Println(float64(negativeCount)/float64(size))
    fmt.Println(float64(zeroesCount)/float64(size))
}

func round(f float64) int {
    return int(f - 0.5)
}
