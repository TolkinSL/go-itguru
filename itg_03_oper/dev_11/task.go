package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var x, y int
	var result bool

	x, y = ReadInput()
	
	if x == y {
		result = true
	} else {
		result = false
	}

	fmt.Println(result)
}

func ReadInput() (int, int) {
	var x, y int
	var val []string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	val = strings.Split(scanner.Text(), " ")

	x, _ = strconv.Atoi(val[0])
	y, _ = strconv.Atoi(val[1])

	return x, y
}