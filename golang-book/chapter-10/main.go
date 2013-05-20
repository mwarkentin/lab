package main

import (
    "fmt"
    "time"
)

func pinger(c chan<- string) {
    for i := 0; ; i++ {
        c <- "ping"
    }
}

func ponger(c chan<- string) {
    for i := 0; ; i++ {
        c <- "pong"
    }
}

func printer(c <-chan string) {
    for {
        fmt.Println(<- c)
        time.Sleep(time.Second * 1)
    }
}

func main() {
    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        for {
            c1 <- "from 1"
            time.Sleep(time.Second * 2)
        }
    }()

    go func() {
        for {
            c2 <- "from 2"
            time.Sleep(time.Second * 3)
        }
    }()

    go func() {
        for {
            select {
            case msg1 := <- c1:
                fmt.Println(msg1)
            case msg2 := <- c2:
                fmt.Println(msg2)
            }
        }
    }()

    var input string
    fmt.Scanln(&input)
}
