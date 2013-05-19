package main

import "fmt"

func average(xs []float64) float64 {
    total := 0.0
    for _, v := range xs {
        total += v
    }
    return total / float64(len(xs))
}

func add(args ...int) int {
    total := 0
    for _, v := range args {
        total += v
    }
    return total
}

func sum(xs []float64) float64 {
    total := 0.0
    for _, v := range xs {
        total += v
    }
    return total
}

func half(num int) (int, bool) {
    half := num / 2
    isEven := num % 2 == 0
    return half, isEven
}

func main() {
    xs := []float64{98,93,77,82,83}
    fmt.Println(average(xs))
    fmt.Println(sum(xs))
    fmt.Println(add(1,2,3))
    fmt.Println(half(1))
}
