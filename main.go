package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var a, b int
	for {
		a = rand.Intn(100) + 1
		b = rand.Intn(100) + 1

		if ageDiff(a, b) == 30 && (a+5 == 2*(b+5)) {
			break
		}
	}

	fmt.Println(a)
	fmt.Println(b)
}

func ageDiff(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
