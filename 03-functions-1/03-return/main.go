package main

import "fmt"

func main() {
	result := Sum(1, 2, 3, 4, 5)
	fmt.Println(result)
}

func Sum(x, y, z, j, k int) int {
	return x + y + z + j + k
}
