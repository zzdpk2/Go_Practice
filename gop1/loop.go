package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBinary(num int) string {
	result := ""
	for {
		if num == 0 {
			break
		}
		lsb := num % 2
		result = strconv.Itoa(lsb) + result
		num /= 2
	}
	return result
}

func printFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	fmt.Println(
		convertToBinary(2),
		convertToBinary(5),
		convertToBinary(13),
	)
	printFile("abc.txt")
}
