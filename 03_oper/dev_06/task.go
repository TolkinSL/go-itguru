package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var result int

	minutes := ReadInput()
	result = minutes * 60

	fmt.Println(result)
}

func ReadInput() int {
	var minutes int
	
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	minutes, _ = strconv.Atoi(scanner.Text())

	return minutes
}