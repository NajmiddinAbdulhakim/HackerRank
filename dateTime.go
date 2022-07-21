package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		arr := strings.Split(strings.Replace(strings.TrimRight(scanner.Text(), "\n"), ".","-",-1),",")
		fmt.Println(arr[0],arr[1])
		d1, _ := time.Parse("12.03.2018 14:00:15", arr[0])
		d2, _ := time.Parse("12.03.2018 14:00:15", arr[1])
		fmt.Println(d1, d2)
		// t1 := time.Since(d1)
		// fmt.Println(time.Until(d2).Round(t1))
	}
}
