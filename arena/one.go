package main

import (
	"math/rand"
)

func returnMove1(sliceMain [][]int) []int {
	var a, b int
	if move > 30 {
		return []int{-1,-1}
	}
	for {
		a = rand.Intn(19)
		b = rand.Intn(19)
		if sliceMain[a][b] == 0 {
			return []int{a,b}
		}
	}
}
