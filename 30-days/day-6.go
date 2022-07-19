package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    //Enter your code here. Read input from STDIN. Print output to STDOUT
    var T int
    fmt.Scan(&T)
    var arr []string
    scanner := bufio.NewScanner(os.Stdin)
    for i := 0; i < T; i++ {
        var l string
        if scanner.Scan() {
            l = strings.TrimRight(scanner.Text(), "\r\n")
        }
        arr = append(arr, l)
    }
    for _, val := range arr {
        var s1, s2 string
        for i, v := range val {
            if i%2 == 0 {
                s1 += string(v)
            } else if i%2 == 1 {
                s2 += string(v)
            }

        }
        fmt.Println(s1,s2)


    }
}
