package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var n int
	fmt.Scan(&n)
	ma := make(map[string]string)
	var arr []string

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		if scanner.Scan() {
			arr = append(arr, strings.Fields(scanner.Text())...)
			ma[arr[0]] = arr[1]
			arr = nil 
		}
	}
	for i := 0; i < n; i++ {
		if scanner.Scan() {
			arr = append(arr, scanner.Text())
		}
	}
	fmt.Println(arr)
	for _,val := range arr {
		if valeus,ok := ma[val]; ok {
			fmt.Printf("%v=%v\n",val,valeus)
		}else{
			fmt.Println("Not found")
		}
	}
}