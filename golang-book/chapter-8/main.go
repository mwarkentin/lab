package main

import "fmt"

func zero(xPtr *int) {
    *xPtr = 0
}

func one(xPtr *int) {
    *xPtr = 1
}

func main() {
    xPtr := new(int)
    zero(xPtr)
    fmt.Println(*xPtr)
    one(xPtr)
    fmt.Println(*xPtr)
}
