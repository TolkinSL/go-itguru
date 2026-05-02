package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	n := ReadInput()
	var result string

	if n > 0 {
		result = fmt.Sprintf("Число позитивное")
	} else if n == 0 {
		result = fmt.Sprintf("Число равно 0")
	} else {
		result = fmt.Sprintf("Число негативное")
	}

	fmt.Println(result)
}

func ReadInput() int {
	var value int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	value, _ = strconv.Atoi(scanner.Text())

	return value
}
