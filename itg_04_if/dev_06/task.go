package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var result string
	x, y, direction := ReadInput()

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

	x = clamp(x, 0, 100)
	y = clamp(y, 0, 100)

	result = fmt.Sprintf("x: %d, y: %d, direction: %s", x, y, direction)
	fmt.Println(result)
}

func clamp(val, min, max int) int {
	if val < min {
		return min
	}
	
	if val > max {
		return max
	} 
	
	return val
}

func ReadInput() (int, int, string) {
	var x, y int
	var direction string
	var val []string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	val = strings.Split(scanner.Text(), " ")

	x, _ = strconv.Atoi(val[0])
	y, _ = strconv.Atoi(val[1])
	direction = val[2]

	return x, y, direction
}