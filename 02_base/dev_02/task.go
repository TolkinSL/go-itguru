package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	msg := scanner.Text()

	return msg
}

func main() {
	myInput := ReadInput()
	fmt.Println(myInput)
}
