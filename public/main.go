package main

import (
	"go-wasm/pkg/utils"
)

func main() {
	println("adding two numbers:", add(2, 3)) // expecting 5
}

// call js methed
func add(x, y int) int

// export golang method
//export multiply
func multiply(x, y int) int {
	return utils.Multiply(x, y)
}
