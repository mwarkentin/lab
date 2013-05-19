package main

import "fmt"

func zero(xPtr *int) {
    *xPtr = 0
}

func one(xPtr *int) {
    *xPtr = 1
}

func square(x *float64) {
    *x = *x * *x
}

func swap(xPtr, yPtr *int) {
    tmp := *xPtr
    *xPtr = *yPtr
    *yPtr = tmp
}

func main() {
    xPtr := new(int)
    zero(xPtr)
    fmt.Println(*xPtr)
    one(xPtr)
    fmt.Println(*xPtr)

    x := 1.5
    square(&x)
    fmt.Println(x)

    y := 1
    z := 2
    swap(&y, &z)
    fmt.Println(y, z)
}
