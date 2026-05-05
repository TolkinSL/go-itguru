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

	result = 1
	
	for i := 1; i <= num; i++ {
		if i%2 != 0 {
			result *= i
		}
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