package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	var result string
	str := ReadInput()

	// myRune := []rune(str)
	runeCount := utf8.RuneCountInString(str)

	// for i := 0; i < len(myRune); i++ {
	// 	result += "*"
	// } 

	for i := 0; i < runeCount; i++ {
		result += "*"
	}

	// fmt.Println(strings.Repeat("*", len([]rune(message))))
	
	fmt.Println(result)
}

func ReadInput() string{
	var message string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	message = scanner.Text()

	return message
}