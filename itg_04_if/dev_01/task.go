package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	n := ReadInput()
	var result bool

	result = (n % 2 !=0) && (n % 7 == 0)
	fmt.Println(result)
}

func ReadInput() int {
	var n int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	
	return n
}