package main

import (
	"fmt"
	"math"
)

func inc(a int8) (int8, error) {
	if a == math.MaxInt8 {
		return 0, fmt.Errorf("error overflow")
	}

	return a + 1, nil
}

func add(a, b int8) (int8, error) {
	fmt.Printf("Res: %d + %d = %d\n", a, b, a + b)
	if a > math.MaxInt8 - b {
		return 0, fmt.Errorf("error overflow plus")
	} else if a < math.MinInt8 - b {
		return 0, fmt.Errorf("error overflow minus")
	}

	return a + b, nil
}

func main() {
	var signed int8 = math.MaxInt8
	signed++

	var unsigned uint8 = math.MaxUint8
	unsigned++

	fmt.Println("signed:", signed)
	fmt.Println("unsigned:", unsigned)

	fmt.Println("inc")
	fmt.Println(inc(126))
	fmt.Println(inc(127))
	fmt.Println(inc(-128))

	fmt.Println("\nadd")
	fmt.Println(add(127, 1))
	fmt.Println(add(127, 0))
	fmt.Println(add(120, 3))
	fmt.Println(add(-10, -120))
}
