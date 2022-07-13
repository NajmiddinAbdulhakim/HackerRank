package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func miniMaxSum(arr []int32) {
	// Write your code here
	// my code
	var ara []int64
	var val int64
	for i, _ := range arr {
		for ind, v := range arr {
			if ind != i {
				val = val + int64(v)
			}
		}
		ara = append(ara, val)
		val = 0
	}
	min := ara[0]
	max := ara[0]
	for _, v := range ara {
		if max < v {
			max = v
		}
		if min > v {
			min = v
		}
	}
	fmt.Println(min, max)
	// my code
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
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
