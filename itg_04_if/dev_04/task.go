package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	stars := ReadInput()
	var result string

	for i := 0; i < stars; i++ {
		result += "★"
	}

	fmt.Println(result)
}

func ReadInput() int {
	var stars int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	stars, _ = strconv.Atoi(scanner.Text())

	return stars
}
