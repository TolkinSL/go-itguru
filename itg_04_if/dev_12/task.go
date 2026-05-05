package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var result int
	num := ReadInput()

	if num > 0 {
		num *= 2
	}

	result = num
	fmt.Println(result)
}

func ReadInput() int {
	var num int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	num, _ = strconv.Atoi(scanner.Text())

	return num
}