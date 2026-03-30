// go build -gcflags="-m -l" main.go

package main

func stackCase() int {
	x := 10
	y := x + 1
	return y
}

func heapCase() *int {
	x := 10
	return &x
}

func main() {
	a := stackCase()
	b := heapCase()

	_ = a
	_ = b
}