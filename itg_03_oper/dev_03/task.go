package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	x, y := ReadInput()
	result := x - y

	fmt.Println(result)
}

func ReadInput() (int, int) {
	var x, y int
	var str string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str = scanner.Text()

	val := strings.Split(str, " ")
	x, _ = strconv.Atoi(val[0])
	y, _ = strconv.Atoi(val[1])

	return x, y
}
