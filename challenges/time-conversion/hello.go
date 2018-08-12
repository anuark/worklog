// https://www.hackerrank.com/challenges/time-conversion/problem
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
    "time"
)

var _ = math.Pi
var _ = bufio.MaxScanTokenSize
var _ = os.DevNull
var _ = strings.Compare
var _ = strconv.IntSize
var _ = log.LUTC
var _ = fmt.Errorf
var _ = flag.ContinueOnError

const longKitchen = "03:04:05PM"
const hour24Format = "15:04:05"

func main() {
    reader := bufio.NewReader(os.Stdin)
    inputTime, _, _ := reader.ReadLine()

    t, _ := time.Parse(longKitchen, string(inputTime))

    fmt.Println(t.Format(hour24Format))
}
