//Запуск Stepik заданий
//echo "Eyjafjallajokull 6" | go run ./shablon.go
//echo "12345 5" | go run ./shablon.go

package main

import (
    "fmt"
    "strings"
    "bufio"
    "os"
    "strconv"
)

func sum(x, y int) int {
	return x + y
}

func main() {
    x, y := ReadInput()
    result := sum(x, y)
    fmt.Println(result)
}

func ReadInput() (int, int) {
    var x, y int
    var input string 
    var values []string
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    input = scanner.Text()
    values = strings.Split(input, " ")
    x, _ = strconv.Atoi(values[0])
    y, _ = strconv.Atoi(values[1])
    return x, y
}