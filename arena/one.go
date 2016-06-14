package main

import (
	"math/rand"
	"fmt"
)

func returnMove1(sliceMain [][]int) []int {
	var a, b int
	for {
		a = rand.Intn(19)
		b = rand.Intn(19)
		fmt.Println("BOT moved:",b,a)
		if sliceMain[a][b] == 0 {
			return []int{a,b}
		}
	}
}
