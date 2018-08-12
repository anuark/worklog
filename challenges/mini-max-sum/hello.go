// https://www.hackerrank.com/challenges/mini-max-sum
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
    stdin := scanner.Text()
    splitStdin := strings.Split(stdin, " ")
    intStdin := make(IntSlice, 0, 5)
    for i := 0; i < len(splitStdin); i++ {
        n, _ := strconv.Atoi(splitStdin[i])
        intStdin = append(intStdin, int64(n))
    }

    allSums := make([]int64, 0, 5)
    for i := 0; i < 5; i++ {
        tmpSlice := make(IntSlice, 0)
        tmpSlice = append(tmpSlice, intStdin...) // copying
        tmpSlice.Remove(i)
        s := sum(tmpSlice...)
        allSums = append(allSums, s)
    }

    var max float64
    for i := 0; i < 5; i++ {
        max = math.Max(max, float64(allSums[i]))
    }

    min := max
    for i := 0; i < 5; i++ {
        min = math.Min(min, float64(allSums[i]))
    }

    fmt.Printf("%0.f %0.f", min, max)
}

type IntSlice []int64

func (p *IntSlice) Remove(i int) {
    s := *p
    s = append(s[:i], s[i+1:]...)
    *p = s
}

func sum(nums ...int64) int64 {
    total := int64(0)

    for _, num := range nums {
        total += num
    }

    return total
}
