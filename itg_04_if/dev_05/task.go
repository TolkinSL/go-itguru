package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	x, y, direction := ReadInput()
	var result string

	switch direction {
	case "up":
		y -= 1
	case "down":
		y += 1
	case "left":
		x -= 1
	case "right":
		x += 1
	}

	result = fmt.Sprintf("x: %d, y: %d, direction: %s", x, y, direction)
	fmt.Println(result)
}

func ReadInput() (int, int, string) {
	var x, y int
	var (
		input     []string
		direction string
	)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = strings.Fields(scanner.Text())
	x, _ = strconv.Atoi(input[0])
	y, _ = strconv.Atoi(input[1])
	direction = input[2]

	return x, y, direction
}
