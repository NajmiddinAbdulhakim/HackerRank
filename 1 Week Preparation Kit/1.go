package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)
// my code
func plusMinus(arr []int32) {
	// Write your code here
	n := len(arr)
	var positive, negative, zero float32
	for _, num := range arr {
		switch {
		case num > 0:
			positive++
		case num < 0:
			negative++
		case num == 0:
			zero++

		}
	}
	fmt.Printf("%f\n", positive/float32(n))
	fmt.Printf("%f\n", negative/float32(n))
	fmt.Printf("%f\n", zero/float32(n))
}
// my code


// No my code
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
