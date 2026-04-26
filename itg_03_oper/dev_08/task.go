package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var num int
	var result string

	num = ReadInput()
	if num % 2 == 0 {
		result =  "четное"
	} else {
		result = "нечетное"
	}

	fmt.Println(result)
}

func ReadInput() int {
	var num int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	num, _ = strconv.Atoi(scanner.Text())

	return num
}