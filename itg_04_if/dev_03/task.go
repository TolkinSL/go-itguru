package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func minmax(arr ...int) string {
	min, max := arr[0], arr[0]

	for _, val := range arr {
		if val > max {
			max = val
		} else if val < min {
			min = val
		}
	}

	return fmt.Sprintf("минимальное: %d, максимальное: %d", min, max)

	// Sort
	// sort.Ints(t)
    // result = fmt.Sprintf("минимальное: %d, максимальное: %d", t[0], t[len(t)-1])
}

func main() {
	x1, x2, x3 := ReadInput()
	var result string

	result = minmax(x1, x2, x3)
	fmt.Println(result)
}

func ReadInput() (int, int, int) {
	var x1, x2, x3 int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	
	input := scanner.Text()
	val := strings.Split(input, " ")
	x1, _ = strconv.Atoi(val[0])
	x2, _ = strconv.Atoi(val[1])
	x3, _ = strconv.Atoi(val[2])

	return x1, x2, x3
}
