// https://www.hackerrank.com/challenges/birthday-cake-candles/problem
package main

import (
    "math"
    "fmt"
    "log"
    "bufio"
    "os"
    "strings"
    "strconv"
    "flag"
)

var _ = math.Pi
var _ = bufio.MaxScanTokenSize
var _ = os.DevNull
var _ = strings.Compare
var _ = strconv.IntSize
var _ = log.LUTC
var _ = fmt.Errorf
var _ = flag.ContinueOnError

// const maxCap = 500*1024
const maxCap = 1000*1024

func main() {

    scanner := bufio.NewScanner(os.Stdin)
    buf := make([]byte, maxCap)
    scanner.Buffer(buf, maxCap)
    scanner.Scan() // number of candles
    scanner.Scan() // candles height 
    input := strings.Split(scanner.Text(), " ")

    candleHeights := make([]int, 0)
    var biggestCandle float64
    for _, v := range input {
        height, _ := strconv.Atoi(v)
        biggestCandle = math.Max(biggestCandle, float64(height))
        candleHeights = append(candleHeights, height)
    }

    candlesHeightsMap := make(map[int]int)
    for _, v := range candleHeights {
        candlesHeightsMap[v]++
    }

    fmt.Println(candlesHeightsMap[int(biggestCandle)])
}
