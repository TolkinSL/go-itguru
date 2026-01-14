package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var role int
	var result string

	role = ReadInput()
	switch role {
	case 1:
		result = "admin"
	case 2:
		result = "moderator"
	case 3:
		result = "user"
	default:
		result = "guest"
	}

	fmt.Println(result)
}

func ReadInput() int {
	var role int
	var val string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	val = scanner.Text()

	role, _ = strconv.Atoi(val)
	return role
}