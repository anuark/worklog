// https://www.hackerrank.com/challenges/staircase/problem
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
var _ = fmt.Errorf


func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    inputN := scanner.Text()
    n, _ := strconv.Atoi(inputN)

    for i := 0; i < n; i++ {
        j := i
        out := "#"
        for o := 0; o < n-1; o++ {
            if j > 0 {
                out = "#" + out
                j--
            } else {
                out = " " + out
            }
        }

        fmt.Println(out)
    }
}
