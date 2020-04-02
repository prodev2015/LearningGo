package main

import (
	"fmt"
)

func add(x, y float32) float32 {
	return x + y
}

func multiple(str1, str2 string) (string, string) {
	return str1, str2
}

func main() {
	var num1, num2 float32 = 5.2, 6.3
	str1 := "sss"
	fmt.Println(str1)
	fmt.Println(add(num1, num2))
	fmt.Println(multiple("hey", "there"))
}
