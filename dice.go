package main

import (
	"fmt"
	"math/rand"
)

func dice() int {
	fmt.Println("Rolling dice...")
	min, max, sum := 1, 6, 0

	for i := 0; i < 2; i++ {
		result := rand.Intn(max-min) + min // minとmaxが範囲となる
		sum += result
		fmt.Printf("Die %d: %d\n", i, result)
	}
	return sum
}

func main() {
	result := dice()
	fmt.Printf("Total value: %d\n", result)
	if result > 7 {
		fmt.Println("You won")
	} else {
		fmt.Println("You lost")
	}
}
