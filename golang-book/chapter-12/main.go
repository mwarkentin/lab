package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

func main() {
    fmt.Println(
        // true
        strings.Contains("test", "es"),

        // 2
        strings.Count("test", "t"),

        // true
        strings.HasPrefix("test", "te"),

        // true
        strings.HasSuffix("test", "st"),

        // 1
        strings.Index("test", "e"),

        // "a-b"
        strings.Join([]string{"a","b"}, "-"),

        // == "aaaaa"
        strings.Repeat("a", 5),

        // "bbaa"
        strings.Replace("aaaa", "a", "b", 2),

        // []string{"a","b","c","d","e"}
        strings.Split("a-b-c-d-e", "-"),

        // "test"
        strings.ToLower("TEST"),

        // "TEST"
        strings.ToUpper("test"),
    )

    arr := []byte("test")
    fmt.Println(arr)

    str := string([]byte{'t','e','s','t'})
    fmt.Println(str)

    bs, err := ioutil.ReadFile("test.txt")
    if err != nil {
        return
    }
    text := string(bs)
    fmt.Println(text)

    file, err := os.Create("test2.txt")
    if err != nil {
        // handle the error here
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close()
    file.WriteString("test")

    dir, err := os.Open(".")
    if err != nil {
        return
    }
    defer dir.Close()

    fileInfos, err := dir.Readdir(-1)
    if err != nil {
        return
    }
    for _, fi := range fileInfos {
        fmt.Println(fi.Name())
    }
}
