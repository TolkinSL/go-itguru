package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var result int
	n := ReadInput()

	for i := 1; i <= n; i++ {
		if i%2 != 0 {
			result += i
		}
	}

	fmt.Println(result)
}

func ReadInput() int {
	var val int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	val, _ = strconv.Atoi(scanner.Text())

	return val
}