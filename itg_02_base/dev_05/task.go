package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	str, score, bonus := ReadInput()
	result := fmt.Sprintf("%s %d", str, score * bonus)
	fmt.Println(result)
}

func ReadInput() (string, int, int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	arr := strings.Split(scanner.Text(), " | ")
	str := arr[0]
	num1, _ := strconv.Atoi(arr[1])
	num2, _ := strconv.Atoi(arr[2])

	return str, num1, num2
}
