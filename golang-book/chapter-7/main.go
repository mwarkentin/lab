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

func max(args ...int) int {
    max := 0
    for _, v := range args {
        if v > max {
            max = v
        }
    }
    return max
}

func makeOddGenerator() func() uint {
    i := uint(1)
    return func() (ret uint) {
        ret = i
        i += 2
        return
    }
}

func fib(x uint) uint {
    if x == 0 || x == 1 {
        return x
    }
    return fib(x-1) + fib(x-2)
}

func main() {
    xs := []float64{98,93,77,82,83}
    fmt.Println(average(xs))
    fmt.Println(sum(xs))
    fmt.Println(add(1,2,3))
    fmt.Println(half(1))
    fmt.Println(max(1,5,23,8,15))

    nextOdd := makeOddGenerator()
    fmt.Println(nextOdd())
    fmt.Println(nextOdd())
    fmt.Println(nextOdd())
    fmt.Println(nextOdd())
    fmt.Println(nextOdd())

    fmt.Println(fib(10))
}
