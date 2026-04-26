package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var hours, result int
	hours = ReadInput()

	result = hours * 3600

	fmt.Println(result)
}

func ReadInput() int {
	var hours int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	hours, _ = strconv.Atoi(scanner.Text())

	return hours
}
