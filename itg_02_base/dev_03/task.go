package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	msg1, msg2 := ReadInput()
	fmt.Println(msg1, msg2)
}

func ReadInput() (string, string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	arr := strings.Split(scanner.Text(), " | ")

	return arr[0], arr[1]
}