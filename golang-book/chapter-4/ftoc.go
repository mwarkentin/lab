package main

import "fmt"

func main() {
    fmt.Println("Enter a temperature (F): ")
    var degreesFarenheight float64
    fmt.Scanf("%f", &degreesFarenheight)

    degreesCelsius := (degreesFarenheight - 32) * 5 / 9

    fmt.Println(degreesCelsius)
}
