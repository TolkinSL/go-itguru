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
	var op string
	var result int

	x, y, op = ReadInput()

	if y == 0 && op == "/" {
		fmt.Println("Divide 0")
		return
	}

	switch op {
	case "x":
		result = x * y
	case "/":
		result = x / y
	case "":
		result = 0
	}

	fmt.Println(result)
}

func ReadInput() (int, int, string) {
	var x, y int
	var op string
	var val []string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	val = strings.Split(scanner.Text(), " ")

	x, _ = strconv.Atoi(val[0])
	op = val[1]
	y, _ = strconv.Atoi(val[2])

	return x, y, op
}
