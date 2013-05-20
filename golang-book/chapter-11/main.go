package main

import (
    "fmt"
    "golang-book/chapter-11/math"
)

func main() {
    xs := []float64{1,2,3,4}
    avg := math.Average(xs)
    fmt.Println(avg)
}
