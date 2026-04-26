package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sum(x, y int) int {
	return x + y
}

func main() {
	x, y := ReadInput()
	result := sum(x, y)
	fmt.Println(result)
}

func ReadInput() (int, int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	val := strings.Split(str, " ")
	x, _ := strconv.Atoi(val[0])
	y, _ := strconv.Atoi(val[1])

	return x, y
}
