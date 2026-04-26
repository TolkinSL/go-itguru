package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var age int
	var result string

	age = ReadInput()

	if age >= 18 {
		result = "взрослый"
	} else {
		result = "подросток"
	}

	fmt.Println(result)
}

func ReadInput() int {
	var age int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	age, _ = strconv.Atoi(scanner.Text())
	
	return age	
}