package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	myStr, myNum := ReadInput()
	fmt.Println(myStr, myNum * 2)
}

func ReadInput() (string, int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	str := strings.Split(scanner.Text(), " | ")
	myStr := str[0]
	myNum, _ := strconv.Atoi(str[1])

	return myStr, myNum
}